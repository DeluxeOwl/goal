package goal

import "fmt"

func (ctx *Context) execute(ops []opcode) (int, error) {
	for ip := 0; ip < len(ops); {
		op := ops[ip]
		//fmt.Printf("op: %s\n", op)
		ip++
		switch op {
		case opNop:
		case opConst:
			ctx.push(ctx.constants[ops[ip]])
			ip++
		case opNil:
			ctx.push(V{})
		case opGlobal:
			x := ctx.globals[ops[ip]]
			if x == (V{}) {
				return ip - 1, fmt.Errorf("undefined global: %s",
					ctx.gNames[ops[ip]])
			}
			ctx.push(x)
			ip++
		case opLocal:
			x := ctx.stack[ctx.frameIdx-int32(ops[ip])]
			if x == (V{}) {
				return ip - 1, fmt.Errorf("undefined local: %s",
					ctx.lambdas[ctx.lambda].Names[int32(ops[ip])])
			}
			ctx.push(x)
			ip++
		case opAssignGlobal:
			ctx.globals[ops[ip]] = ctx.top()
			ip++
		case opAssignLocal:
			ctx.stack[ctx.frameIdx-int32(ops[ip])] = ctx.top()
			ip++
		case opVariadic:
			ctx.push(newBV(Variadic(ops[ip])))
			ip++
		case opLambda:
			ctx.push(ctx.lambdaVs[ops[ip]])
			ip++
		case opApply:
			err := ctx.popApplyN(1)
			if err != nil {
				return ip - 1, err
			}
		case opApplyV:
			v := Variadic(ops[ip])
			r := ctx.applyVariadic(v)
			if isErr(r) {
				return ip - 1, r.BV.(error)
			}
			ctx.push(r)
			ip++
		case opApply2:
			err := ctx.popApplyN(2)
			if err != nil {
				return ip - 1, err
			}
		case opApply2V:
			v := Variadic(ops[ip])
			r := ctx.applyNVariadic(v, 2)
			if isErr(r) {
				return ip - 1, r.BV.(error)
			}
			ctx.push(r)
			ip++
		case opApplyN:
			err := ctx.popApplyN(int(ops[ip]))
			if err != nil {
				return ip - 1, err
			}
			ip++
		case opApplyNV:
			v := Variadic(ops[ip])
			ip++
			r := ctx.applyNVariadic(v, int(ops[ip]))
			if isErr(r) {
				return ip - 2, r.BV.(error)
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
		}
		//fmt.Printf("stack: %v\n", ctx.stack)
	}
	return len(ops), nil
}

func (ctx *Context) popApplyN(n int) error {
	x := ctx.pop()
	r := ctx.applyN(x, n)
	if isErr(r) {
		return r.BV.(error)
	}
	ctx.push(r)
	return nil
}

const maxCallDepth = 100000

func (ctx *Context) push(x V) {
	ctx.stack = append(ctx.stack, x)
}

func (ctx *Context) pushArgs(args []V) {
	ctx.stack = append(ctx.stack, args...)
}

func (ctx *Context) pop() V {
	arg := ctx.stack[len(ctx.stack)-1]
	ctx.stack[len(ctx.stack)-1] = V{}
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
	return arg
}

func (ctx *Context) top() V {
	return ctx.stack[len(ctx.stack)-1]
}

func (ctx *Context) popN(n int) []V {
	topN := ctx.stack[len(ctx.stack)-n:]
	args := cloneArgs(topN)
	for i := range topN {
		topN[i] = V{}
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
	ctx.stack[len(ctx.stack)-1] = V{}
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
}

func (ctx *Context) dropN(n int) {
	topN := ctx.stack[len(ctx.stack)-n:]
	for i := range topN {
		topN[i] = V{}
	}
	ctx.stack = ctx.stack[:len(ctx.stack)-n]
}
