package goal

import "unicode/utf8"

// VUTF8RCount implements the "utf8.rcount" variadic verb.
func VUTF8RCount(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		return utf8rcount(args[0])
	default:
		return panicRank("utf8.rcount")
	}
}

func utf8rcount(x V) V {
	switch xv := x.value.(type) {
	case S:
		return NewI(int64(utf8.RuneCountInString(string(xv))))
	case *AS:
		r := make([]int64, xv.Len())
		for i, s := range xv.Slice {
			r[i] = int64(utf8.RuneCountInString(s))
		}
		return NewAI(r)
	case *AV:
		r := make([]V, xv.Len())
		for i, xi := range xv.Slice {
			r[i] = utf8rcount(xi)
			if r[i].IsPanic() {
				return r[i]
			}
		}
		return Canonical(NewAV(r))
	default:
		return panicType("utf8.rcount x", "x", x)
	}
}

// VUTF8Valid implements the "utf8.valid" variadic verb.
func VUTF8Valid(ctx *Context, args []V) V {
	switch len(args) {
	case 1:
		return utf8valid(args[0])
	default:
		return panicRank("utf8.valid")
	}
}

func utf8valid(x V) V {
	switch xv := x.value.(type) {
	case S:
		return NewI(b2i(utf8.ValidString(string(xv))))
	case *AS:
		r := make([]bool, xv.Len())
		for i, s := range xv.Slice {
			r[i] = utf8.ValidString(s)
		}
		return NewAB(r)
	case *AV:
		r := make([]V, xv.Len())
		for i, xi := range xv.Slice {
			r[i] = utf8valid(xi)
			if r[i].IsPanic() {
				return r[i]
			}
		}
		return Canonical(NewAV(r))
	default:
		return panicType("utf8.rcount x", "x", x)
	}
}
