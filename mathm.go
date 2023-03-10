// Code generated by scripts/math.goal. DO NOT EDIT.

package goal

import "math"

func mathm(x V, f func(float64) float64) V {
	if x.IsI() {
		return NewF(f(float64(x.I())))
	}
	if x.IsF() {
		return NewF(f(x.F()))
	}
	switch xv := x.value.(type) {
	case *AB:
		return mathm(fromABtoAF(xv), f)
	case *AI:
		return mathm(toAF(xv), f)
	case *AF:
		r := xv.reuse()
		for i, xi := range xv.Slice {
			r.Slice[i] = f(xi)
		}
		return NewV(r)
	case *AV:
		r := xv.reuse()
		for i, xi := range xv.Slice {
			ri := mathm(xi, f)
			if ri.IsPanic() {
				return ri
			}
			r.Slice[i] = ri
		}
		return NewV(r)
	default:
		return Panicf("bad type in x (%s)", x.Type())
	}
}

// VAcos implements the acos variadic.
func VAcos(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Acos)
		if r.IsPanic() {
			return NewPanic("acos x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("acos")
	}
}

// VAsin implements the asin variadic.
func VAsin(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Asin)
		if r.IsPanic() {
			return NewPanic("asin x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("asin")
	}
}

// VAtan implements the atan variadic.
func VAtan(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Atan)
		if r.IsPanic() {
			return NewPanic("atan x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("atan")
	}
}

// VCos implements the cos variadic.
func VCos(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Cos)
		if r.IsPanic() {
			return NewPanic("cos x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("cos")
	}
}

// VExp implements the exp variadic.
func VExp(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Exp)
		if r.IsPanic() {
			return NewPanic("exp x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("exp")
	}
}

// VLog implements the log variadic.
func VLog(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Log)
		if r.IsPanic() {
			return NewPanic("log x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("log")
	}
}

// VRoundToEven implements the round variadic.
func VRoundToEven(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.RoundToEven)
		if r.IsPanic() {
			return NewPanic("round x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("round")
	}
}

// VSin implements the sin variadic.
func VSin(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Sin)
		if r.IsPanic() {
			return NewPanic("sin x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("sin")
	}
}

// VSqrt implements the sqrt variadic.
func VSqrt(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Sqrt)
		if r.IsPanic() {
			return NewPanic("sqrt x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("sqrt")
	}
}

// VTan implements the tan variadic.
func VTan(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		r := mathm(args[0], math.Tan)
		if r.IsPanic() {
			return NewPanic("tan x : " + string(r.value.(panicV)))
		}
		return r
	default:
		return panicRank("tan")
	}
}
