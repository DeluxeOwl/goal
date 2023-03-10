package goal

import "math"

// VNaN implements the nan variadic verb.
func VNaN(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		return isNaN(args[0])
	case 2:
		return fillNaN(args[1], args[0])
	default:
		return panicRank("nan")
	}
}

func isNaN(x V) V {
	if x.IsI() {
		return NewI(b2i(false))
	}
	if x.IsF() {
		return NewF(b2f(math.IsNaN(x.F())))
	}
	switch xv := x.value.(type) {
	case *AB:
		r := make([]bool, xv.Len())
		return NewAB(r)
	case *AI:
		r := make([]bool, xv.Len())
		return NewAB(r)
	case *AF:
		r := make([]bool, xv.Len())
		for i, xi := range xv.Slice {
			r[i] = math.IsNaN(xi)
		}
		return NewAB(r)
	case *AV:
		r := xv.reuse()
		for i, xi := range xv.Slice {
			ri := isNaN(xi)
			if ri.IsPanic() {
				return ri
			}
			r.Slice[i] = ri
		}
		return NewV(r)
	default:
		return panicType("NaN x", "x", x)
	}
}

func fillNaN(x V, y V) V {
	var fill float64
	if x.IsI() {
		fill = float64(x.I())
	} else if x.IsF() {
		fill = x.F()
	} else {
		return panicType("x NaN y", "x", x)
	}
	return fillNaNf(fill, y)
}

func fillNaNf(fill float64, y V) V {
	if y.IsI() {
		return y
	}
	if y.IsF() {
		if math.IsNaN(y.F()) {
			return NewF(fill)
		}
		return y
	}
	switch yv := y.value.(type) {
	case *AB:
		return y
	case *AI:
		return y
	case *AF:
		var r []float64
		if yv.reusable() {
			r = yv.Slice
		} else {
			r = make([]float64, yv.Len())
			copy(r, yv.Slice)
		}
		for i, ri := range r {
			if math.IsNaN(ri) {
				r[i] = fill
			}
		}
		return NewAF(r)
	case *AV:
		r := yv.reuse()
		for i, yi := range yv.Slice {
			ri := fillNaNf(fill, yi)
			if ri.IsPanic() {
				return ri
			}
			r.Slice[i] = ri
		}
		return NewV(r)
	default:
		return panicType("x NaN y", "y", y)
	}
}
