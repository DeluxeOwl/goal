package goal

import "fmt"

func (ctx *Context) execute(ops []opcode) (int, error) {
	for ip := 0; ip < len(ops); {
		op := ops[ip]
		//fmt.Printf("op: %s\n", op)
		ip++
		switch op {
		case opConst:
			ctx.push(ctx.constants[ops[ip]])
			ip++
		case opInt:
			ctx.pushNoRC(NewI(int64(ops[ip])))
			ip++
		case opNil:
			ctx.pushNoRC(V{})
		case opGlobal:
			x := ctx.globals[ops[ip]]
			if x.kind == valNil {
				return ip - 1, fmt.Errorf("undefined global: %s",
					ctx.gNames[ops[ip]])
			}
			ctx.push(x)
			ip++
		case opGlobalLast:
			x := ctx.globals[ops[ip]]
			if x.kind == valNil {
				return ip - 1, fmt.Errorf("undefined global: %s",
					ctx.gNames[ops[ip]])
			}
			ctx.pushNoRC(x)
			ip++
		case opLocal:
			x := ctx.stack[ctx.frameIdx-int32(ops[ip])]
			if x.kind == valNil {
				return ip - 1, fmt.Errorf("undefined local: %s",
					ctx.lambdas[ctx.lambda].Names[int32(ops[ip])])
			}
			ctx.push(x)
			ip++
		case opLocalLast:
			x := ctx.stack[ctx.frameIdx-int32(ops[ip])]
			if x.kind == valNil {
				return ip - 1, fmt.Errorf("undefined local: %s",
					ctx.lambdas[ctx.lambda].Names[int32(ops[ip])])
			}
			ctx.pushNoRC(x)
			ip++
		case opAssignGlobal:
			x := ctx.top()
			x.IncrRC()
			ctx.globals[ops[ip]] = x
			ip++
		case opAssignLocal:
			x := ctx.top()
			x.IncrRC()
			ctx.stack[ctx.frameIdx-int32(ops[ip])] = x
			ip++
		case opVariadic:
			ctx.pushNoRC(newVariadic(variadic(ops[ip])))
			ip++
		case opLambda:
			ctx.pushNoRC(newLambda(lambda(ops[ip])))
			ip++
		case opApply:
			x := ctx.pop()
			r := ctx.applyN(x, 1)
			if r.IsPanic() {
				return ip - 1, newExecError(r)
			}
			ctx.push(r)
		case opApplyV:
			v := variadic(ops[ip])
			r := ctx.applyVariadic(v)
			if r.IsPanic() {
				return ip - 1, newExecError(r)
			}
			ctx.push(r)
			ip++
		case opDerive:
			v := variadic(ops[ip])
			ctx.push(NewV(&derivedVerb{Fun: v, Arg: ctx.pop()}))
			ip++
		case opApply2:
			x := ctx.pop()
			r := ctx.applyN(x, 2)
			if r.IsPanic() {
				return ip - 1, newExecError(r)
			}
			ctx.push(r)
		case opApply2V:
			v := variadic(ops[ip])
			r := ctx.apply2Variadic(v)
			if r.IsPanic() {
				return ip - 1, newExecError(r)
			}
			ctx.push(r)
			ip++
		case opApplyN:
			x := ctx.pop()
			r := ctx.applyN(x, int(ops[ip]))
			if r.IsPanic() {
				return ip - 1, newExecError(r)
			}
			ctx.push(r)
			ip++
		case opApplyNV:
			v := variadic(ops[ip])
			ip++
			r := ctx.applyNVariadic(v, int(ops[ip]))
			if r.IsPanic() {
				return ip - 2, newExecError(r)
			}
			ctx.push(r)
			ip++
		case opDrop:
			ctx.drop()
		case opJump:
			ip += int(ops[ip])
		case opJumpFalse:
			if isFalse(ctx.top()) {
				ip += int(ops[ip])
			} else {
				ip++
			}
		case opJumpTrue:
			if isTrue(ctx.top()) {
				ip += int(ops[ip])
			} else {
				ip++
			}
		case opReturn:
			return len(ops), nil
		case opTry:
			if ctx.top().IsError() {
				return len(ops), nil
			}
		}
		//fmt.Printf("stack: %v\n", ctx.stack)
	}
	return len(ops), nil
}

const maxCallDepth = 100000

func (ctx *Context) swap() {
	ctx.stack[len(ctx.stack)-1], ctx.stack[len(ctx.stack)-2] = ctx.stack[len(ctx.stack)-2], ctx.stack[len(ctx.stack)-1]
}

func (ctx *Context) push(x V) {
	x.IncrRC()
	ctx.stack = append(ctx.stack, x)
}

func (ctx *Context) pushNoRC(x V) {
	ctx.stack = append(ctx.stack, x)
}

func (ctx *Context) pushArgs(args []V) {
	rcincrArgs(args)
	ctx.stack = append(ctx.stack, args...)
}

func (ctx *Context) pop() V {
	arg := ctx.stack[len(ctx.stack)-1]
	if arg.kind == valBoxed {
		arg.rcdecrRefCounter()
		ctx.stack[len(ctx.stack)-1].value = nil
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
	return arg
}

func (ctx *Context) top() V {
	return ctx.stack[len(ctx.stack)-1]
}

func (ctx *Context) popN(n int) []V {
	topN := ctx.stack[len(ctx.stack)-n:]
	args := cloneArgs(topN)
	for i, v := range topN {
		if v.kind == valBoxed {
			v.rcdecrRefCounter()
			topN[i].value = nil
		}
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-n]
	return args
}

func (ctx *Context) peek() []V {
	return ctx.stack[len(ctx.stack)-1:]
}

func (ctx *Context) peekN(n int) []V {
	return ctx.stack[len(ctx.stack)-n:]
}

func (ctx *Context) drop() {
	if v := ctx.stack[len(ctx.stack)-1]; v.kind == valBoxed {
		v.rcdecrRefCounter()
		//ctx.stack[len(ctx.stack)-1].value = nil
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
}

func (ctx *Context) dropNoRC() {
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
}

func (ctx *Context) drop2() {
	if v := ctx.stack[len(ctx.stack)-2]; v.kind == valBoxed {
		v.rcdecrRefCounter()
		//ctx.stack[len(ctx.stack)-2].value = nil
	}
	last := len(ctx.stack) - 1
	if ctx.stack[last].kind == valBoxed {
		ctx.stack[last].rcdecrRefCounter()
		ctx.stack[last].value = nil
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-2]
}

func (ctx *Context) dropN(n int) {
	topN := ctx.stack[len(ctx.stack)-n:]
	for i, v := range topN {
		if v.kind == valBoxed {
			v.rcdecrRefCounter()
			topN[i].value = nil
		}
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-n]
}

func (ctx *Context) dropNnoRC(n int) {
	topN := ctx.stack[len(ctx.stack)-n:]
	for i, v := range topN {
		if v.kind == valBoxed {
			topN[i].value = nil
		}
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-n]
}

func rcincrArgs(args []V) {
	for _, v := range args {
		v.IncrRC()
	}
}
