package goal

//import "fmt"

// Apply calls a value with a single argument.
func (ctx *Context) Apply(v, x V) V {
	ctx.push(x)
	return ctx.applyN(v, 1)
}

// ApplyN calls a value with one or more arguments. The arguments should be
// provided in reverse order, given the stack-based right to left semantics
// used by the language.
func (ctx *Context) ApplyN(v V, args []V) V {
	if len(args) == 0 {
		panic("ApplyArgs: len(args) should be > 0")
	}
	ctx.pushArgs(args)
	return ctx.applyN(v, len(args))
}

// applyN applies v with the top n arguments in the stack. It consumes the
// arguments, but does not push the result, returing it instead.
func (ctx *Context) applyN(v V, n int) V {
	switch v := v.(type) {
	case Lambda:
		return ctx.applyLambda(v, n)
	case Variadic:
		if n == 1 {
			return ctx.applyVariadic(v)
		}
		return ctx.applyNVariadic(v, n)
	case DerivedVerb:
		ctx.push(v.Arg)
		args := ctx.peekN(n + 1)
		if hasNil(args) {
			return Projection{Fun: v, Args: ctx.popN(n + 1)}
		}
		res := ctx.variadics[v.Fun].Func(ctx, args)
		ctx.dropN(n + 1)
		return res
	case ProjectionOne:
		if n > 1 {
			return errf("too many arguments: got %d, expected 1", n)
		}
		arg := ctx.top()
		if arg == nil {
			ctx.drop()
			return v
		}
		ctx.push(v.Arg)
		return ctx.applyN(v.Fun, 2)
	case Projection:
		return ctx.applyProjection(v, n)
	case Composition:
		res := ctx.applyN(v.Right, n)
		_, ok := res.(error)
		if ok {
			return res
		}
		ctx.push(res)
		return ctx.applyN(v.Left, 1)
	case Array:
		switch n {
		case 1:
			return canonical(ctx.applyArray(v, ctx.pop()))
		default:
			args := ctx.peekN(n)
			res := ctx.applyArrayArgs(v, args)
			ctx.dropN(n)
			return canonical(res)
		}
	default:
		return errf("type %s cannot be applied", v.Type())
	}
}

func (ctx *Context) applyArray(v Array, xv V) V {
	if xv == nil {
		return v
	}
	var res V
	switch x := xv.(type) {
	case F:
		if !isI(x) {
			return errf("x[y] : non-integer index: %g", x)
		}
		i := int(x)
		if i < 0 {
			i = v.Len() + i
		}
		if i < 0 || i >= v.Len() {
			return errf("x[y] : out of bounds index: %d", i)
		}
		res = v.At(i)
	case I:
		i := int(x)
		if i < 0 {
			i = v.Len() + i
		}
		if i < 0 || i >= v.Len() {
			return errf("x[y] : out of bounds index: %d", i)
		}
		res = v.At(i)
	case Array:
		indices := toIndices(xv, v.Len())
		if err, ok := indices.(E); ok {
			return err
		}
		res = v.Select(indices.(AI))
	}
	return res
}

func (ctx *Context) applyArrayArgs(v Array, args []V) V {
	if len(args) == 0 {
		return v
	}
	arg := args[len(args)-1]
	res := ctx.applyArray(v, arg)
	if _, ok := res.(E); ok {
		return res
	}
	args = args[:len(args)-1]
	if len(args) == 0 {
		return res
	}
	switch res := res.(type) {
	case AV:
		for i := range res {
			switch z := res[i].(type) {
			case Array:
				res[i] = ctx.applyArrayArgs(z, args)
				if _, ok := res[i].(E); ok {
					return res
				}
			}
		}
		return res
	case Array:
		if len(args) > 1 {
			return errs("x[y] : out of depth index")
		}
		vres := ctx.applyArray(res, args[len(args)-1])
		if _, ok := vres.(E); ok {
			return vres
		}
		return vres
	default:
		return errs("x[y] : out of depth index")
	}
}

func (ctx *Context) applyVariadic(v Variadic) V {
	args := ctx.peek()
	vv := args[0]
	if ctx.variadics[v].Adverb {
		ctx.drop()
		return DerivedVerb{Fun: v, Arg: vv}
	}
	if vv == nil {
		ctx.drop()
		return Projection{Fun: v, Args: []V{vv}}
	}
	//switch vv := vv.(type) {
	//case Variadic:
	//return Composition{Left: v, Right: vv}
	//default:
	if vv == nil {
		return Projection{Fun: v, Args: []V{vv}}
	}
	res := ctx.variadics[v].Func(ctx, args)
	ctx.drop()
	return res
	//}
}

func (ctx *Context) applyNVariadic(v Variadic, n int) V {
	args := ctx.peekN(n)
	if hasNil(args) {
		if n == 2 && args[1] != nil {
			arg := args[1]
			ctx.dropN(n)
			return ProjectionOne{Fun: v, Arg: arg}
		}
		return Projection{Fun: v, Args: ctx.popN(n)}
	}
	//if n == 2 && !ctx.variadics[v].Adverb {
	//switch arg := args[1].(type) {
	//case Lambda:
	//case Function:
	//res := Composition{
	//Left:  ProjectionOne{Fun: v, Arg: args[0]},
	//Right: arg,
	//}
	//ctx.dropN(2)
	//return res
	//}
	//}
	res := ctx.variadics[v].Func(ctx, args)
	ctx.dropN(n)
	return res
}

func (ctx *Context) applyProjection(v Projection, n int) V {
	args := ctx.peekN(n)
	nNils := countNils(v.Args)
	switch {
	case len(args) > nNils:
		return errs("too many arguments")
	case len(args) == nNils:
		n := 0
		for _, v := range v.Args {
			switch {
			case v != nil:
				ctx.push(v)
			default:
				ctx.push(args[n])
				n++
			}
		}
		res := ctx.applyN(v.Fun, len(v.Args))
		ctx.dropN(len(v.Args))
		ctx.dropN(n)
		return res
	default:
		vargs := cloneArgs(v.Args)
		n := 1
		for i := len(vargs) - 1; i >= 0; i-- {
			if vargs[i] == nil {
				if n > len(args) {
					break
				}
				vargs[i] = args[len(args)-n]
				n++
			}
		}
		ctx.dropN(n)
		return Projection{Fun: v, Args: vargs}
	}
}

func (ctx *Context) applyLambda(id Lambda, n int) V {
	if ctx.callDepth > maxCallDepth {
		return errs("lambda: exceeded maximum call depth")
	}
	lc := ctx.prog.Lambdas[int(id)]
	if lc.Rank < n {
		return errf("lambda: too many arguments: got %d, expected %d", n, lc.Rank)
	} else if lc.Rank > n {
		if lc.Rank == 2 && n == 1 {
			return ProjectionOne{Fun: id, Arg: ctx.pop()}
		}
		return Projection{Fun: id, Args: ctx.popN(n)}
	}
	nVars := len(lc.Names) - lc.Rank
	olen := len(ctx.stack)
	for i := 0; i < nVars; i++ {
		ctx.push(nil)
	}
	oframeIdx := ctx.frameIdx
	ctx.frameIdx = int32(len(ctx.stack) - 1)

	olambda := ctx.lambda
	ctx.lambda = int(id)
	ofname := ctx.fname
	ctx.fname = lc.Filename
	ctx.callDepth++
	ip, err := ctx.execute(lc.Body)
	ctx.callDepth--
	ctx.fname = ofname
	ctx.lambda = olambda

	if err != nil {
		ctx.updateErrPos(ip, id, lc.Filename)
		return E(err.Error())
	}
	var res V
	switch len(ctx.stack) {
	case olen + nVars:
	case olen + nVars + 1:
		res = ctx.stack[len(ctx.stack)-1]
		ctx.drop()
	default:
		ctx.updateErrPos(ip, id, lc.Filename)
		// should not happen
		return errf("lambda %d: bad len %d vs old %d (depth: %d): %v", id, len(ctx.stack), olen, ctx.callDepth, ctx.stack)
	}
	if nVars > 0 {
		ctx.dropN(nVars)
	}
	ctx.dropN(n)
	ctx.frameIdx = oframeIdx
	return res
}

func (x AV) Select(y AI) V {
	res := make(AV, len(y))
	xlen := x.Len()
	for i := range res {
		idx := y[i]
		if idx < 0 {
			idx += xlen
		}
		if idx < 0 || idx >= len(x) {
			return errf("x[y] : index out of bounds: %d (length %d)", y[i], len(x))
		}
		res[i] = x[idx]
	}
	return res
}

func (x AB) Select(y AI) V {
	res := make(AB, len(y))
	xlen := x.Len()
	for i := range res {
		idx := y[i]
		if idx < 0 {
			idx += xlen
		}
		if idx < 0 || idx >= len(x) {
			return errf("x[y] : index out of bounds: %d (length %d)", y[i], len(x))
		}
		res[i] = x[idx]
	}
	return res
}

func (x AI) Select(y AI) V {
	res := make(AI, len(y))
	xlen := x.Len()
	for i := range res {
		idx := y[i]
		if idx < 0 {
			idx += xlen
		}
		if idx < 0 || idx >= len(x) {
			return errf("x[y] : index out of bounds: %d (length %d)", y[i], len(x))
		}
		res[i] = x[idx]
	}
	return res
}

func (x AF) Select(y AI) V {
	res := make(AF, len(y))
	xlen := x.Len()
	for i := range res {
		idx := y[i]
		if idx < 0 {
			idx += xlen
		}
		if idx < 0 || idx >= len(x) {
			return errf("x[y] : index out of bounds: %d (length %d)", y[i], len(x))
		}
		res[i] = x[idx]
	}
	return res
}

func (x AS) Select(y AI) V {
	res := make(AS, len(y))
	xlen := x.Len()
	for i := range res {
		idx := y[i]
		if idx < 0 {
			idx += xlen
		}
		if idx < 0 || idx >= len(x) {
			return errf("x[y] : index out of bounds: %d (length %d)", y[i], len(x))
		}
		res[i] = x[idx]
	}
	return res
}
