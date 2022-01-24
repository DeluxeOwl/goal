package main

import (
	"math"
	"strings"
)

func Equal(w, x O) O {
	switch w := w.(type) {
	case B:
		return EqualBO(w, x)
	case F:
		return EqualFO(w, x)
	case I:
		return EqualIO(w, x)
	case S:
		return EqualSO(w, x)
	case AB:
		return EqualABO(w, x)
	case AF:
		return EqualAFO(w, x)
	case AI:
		return EqualAIO(w, x)
	case AS:
		return EqualASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("=")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Equal(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Equal(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w == x
	case F:
		return B2F(w) == x
	case I:
		return B2I(w) == x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) == x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := EqualBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w == B2F(x)
	case F:
		return w == x
	case I:
		return w == F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := EqualFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w == B2I(x)
	case F:
		return F(w) == x
	case I:
		return w == x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := EqualIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w == x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w == x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := EqualSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == x
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) == x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) == x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) == x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := EqualABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := EqualAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) == x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) == x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := EqualAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func EqualASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] == x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] == x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := EqualASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("=")
	}
}

func NotEqual(w, x O) O {
	switch w := w.(type) {
	case B:
		return NotEqualBO(w, x)
	case F:
		return NotEqualFO(w, x)
	case I:
		return NotEqualIO(w, x)
	case S:
		return NotEqualSO(w, x)
	case AB:
		return NotEqualABO(w, x)
	case AF:
		return NotEqualAFO(w, x)
	case AI:
		return NotEqualAIO(w, x)
	case AS:
		return NotEqualASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("≠")
			}
			r := make(AO, len(w))
			for i := range r {
				v := NotEqual(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := NotEqual(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w != x
	case F:
		return B2F(w) != x
	case I:
		return B2I(w) != x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) != x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := NotEqualBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w != B2F(x)
	case F:
		return w != x
	case I:
		return w != F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := NotEqualFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w != B2I(x)
	case F:
		return F(w) != x
	case I:
		return w != x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := NotEqualIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w != x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w != x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := NotEqualSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != x
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) != x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) != x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) != x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := NotEqualABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := NotEqualAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) != x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) != x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := NotEqualAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func NotEqualASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] != x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] != x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := NotEqualASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≠")
	}
}

func Lesser(w, x O) O {
	switch w := w.(type) {
	case B:
		return LesserBO(w, x)
	case F:
		return LesserFO(w, x)
	case I:
		return LesserIO(w, x)
	case S:
		return LesserSO(w, x)
	case AB:
		return LesserABO(w, x)
	case AF:
		return LesserAFO(w, x)
	case AI:
		return LesserAIO(w, x)
	case AS:
		return LesserASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("<")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Lesser(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Lesser(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return !w && x
	case F:
		return B2F(w) < x
	case I:
		return B2I(w) < x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = !w && x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) < x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w < B2F(x)
	case F:
		return w < x
	case I:
		return w < F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w < B2I(x)
	case F:
		return F(w) < x
	case I:
		return w < x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w < x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w < x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = !w[i] && x
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) < x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) < x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = !w[i] && x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) < x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) < x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) < x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] < x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] < x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("<")
	}
}

func LesserEq(w, x O) O {
	switch w := w.(type) {
	case B:
		return LesserEqBO(w, x)
	case F:
		return LesserEqFO(w, x)
	case I:
		return LesserEqIO(w, x)
	case S:
		return LesserEqSO(w, x)
	case AB:
		return LesserEqABO(w, x)
	case AF:
		return LesserEqAFO(w, x)
	case AI:
		return LesserEqAIO(w, x)
	case AS:
		return LesserEqASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("≤")
			}
			r := make(AO, len(w))
			for i := range r {
				v := LesserEq(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := LesserEq(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return x || !w
	case F:
		return B2F(w) <= x
	case I:
		return B2I(w) <= x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = x[i] || !w
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) <= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserEqBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w <= B2F(x)
	case F:
		return w <= x
	case I:
		return w <= F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserEqFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w <= B2I(x)
	case F:
		return F(w) <= x
	case I:
		return w <= x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserEqIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w <= x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w <= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := LesserEqSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = x || !w[i]
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) <= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) <= x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = x[i] || !w[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) <= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserEqABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserEqAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) <= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) <= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserEqAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func LesserEqASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] <= x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] <= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := LesserEqASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≤")
	}
}

func Greater(w, x O) O {
	switch w := w.(type) {
	case B:
		return GreaterBO(w, x)
	case F:
		return GreaterFO(w, x)
	case I:
		return GreaterIO(w, x)
	case S:
		return GreaterSO(w, x)
	case AB:
		return GreaterABO(w, x)
	case AF:
		return GreaterAFO(w, x)
	case AI:
		return GreaterAIO(w, x)
	case AS:
		return GreaterASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen(">")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Greater(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Greater(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w && !x
	case F:
		return B2F(w) > x
	case I:
		return B2I(w) > x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w && !x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) > x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w > B2F(x)
	case F:
		return w > x
	case I:
		return w > F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w > B2I(x)
	case F:
		return F(w) > x
	case I:
		return w > x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w > x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w > x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] && !x
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) > x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) > x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] && !x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) > x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) > x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) > x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] > x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] > x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype(">")
	}
}

func GreaterEq(w, x O) O {
	switch w := w.(type) {
	case B:
		return GreaterEqBO(w, x)
	case F:
		return GreaterEqFO(w, x)
	case I:
		return GreaterEqIO(w, x)
	case S:
		return GreaterEqSO(w, x)
	case AB:
		return GreaterEqABO(w, x)
	case AF:
		return GreaterEqAFO(w, x)
	case AI:
		return GreaterEqAIO(w, x)
	case AS:
		return GreaterEqASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("≥")
			}
			r := make(AO, len(w))
			for i := range r {
				v := GreaterEq(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := GreaterEq(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w || !x
	case F:
		return B2F(w) >= x
	case I:
		return B2I(w) >= x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w || !x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w) >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w) >= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterEqBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w >= B2F(x)
	case F:
		return w >= x
	case I:
		return w >= F(x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterEqFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w >= B2I(x)
	case F:
		return F(w) >= x
	case I:
		return w >= x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w) >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterEqIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w >= x
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w >= x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := GreaterEqSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] || !x
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2F(w[i]) >= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = B2I(w[i]) >= x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] || !x[i]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2F(w[i]) >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = B2I(w[i]) >= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterEqABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= B2F(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= F(x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= B2F(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterEqAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= B2I(x)
		}
		return r
	case F:
		r := make(AB, len(w))
		for i := range r {
			r[i] = F(w[i]) >= x
		}
		return r
	case I:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= B2I(x[i])
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i := range r {
			r[i] = F(w[i]) >= x[i]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterEqAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func GreaterEqASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] >= x
		}
		return r
	case AS:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] >= x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := GreaterEqASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("≥")
	}
}

func Add(w, x O) O {
	switch w := w.(type) {
	case B:
		return AddBO(w, x)
	case F:
		return AddFO(w, x)
	case I:
		return AddIO(w, x)
	case S:
		return AddSO(w, x)
	case AB:
		return AddABO(w, x)
	case AF:
		return AddAFO(w, x)
	case AI:
		return AddAIO(w, x)
	case AS:
		return AddASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("+")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Add(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Add(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return B2I(w) + B2I(x)
	case F:
		return B2F(w) + x
	case I:
		return B2I(w) + x
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) + B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w) + x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) + x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AddBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w + B2F(x)
	case F:
		return w + x
	case I:
		return w + F(x)
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w + B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w + x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w + F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AddFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w + B2I(x)
	case F:
		return F(w) + x
	case I:
		return w + x
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w + B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w) + x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w + x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AddIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return w + x
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = w + x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AddSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) + B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = B2F(w[i]) + x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) + x
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) + B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w[i]) + x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) + x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AddABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] + B2F(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] + x
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] + F(x)
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] + B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] + x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] + F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AddAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] + B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(w[i]) + x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] + x
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] + B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w[i]) + x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] + x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AddAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func AddASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = w[i] + x
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = w[i] + x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AddASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("+")
	}
}

func Subtract(w, x O) O {
	switch w := w.(type) {
	case B:
		return SubtractBO(w, x)
	case F:
		return SubtractFO(w, x)
	case I:
		return SubtractIO(w, x)
	case AB:
		return SubtractABO(w, x)
	case AF:
		return SubtractAFO(w, x)
	case AI:
		return SubtractAIO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("-")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Subtract(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Subtract(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return B2I(w) - B2I(x)
	case F:
		return B2F(w) - x
	case I:
		return B2I(w) - x
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) - B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w) - x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) - x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := SubtractBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w - B2F(x)
	case F:
		return w - x
	case I:
		return w - F(x)
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w - B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w - x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w - F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := SubtractFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w - B2I(x)
	case F:
		return F(w) - x
	case I:
		return w - x
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w - B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w) - x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w - x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := SubtractIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) - B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = B2F(w[i]) - x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) - x
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) - B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w[i]) - x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) - x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := SubtractABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] - B2F(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] - x
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] - F(x)
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] - B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] - x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] - F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := SubtractAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func SubtractAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] - B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(w[i]) - x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] - x
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] - B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w[i]) - x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] - x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := SubtractAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("-")
	}
}

func Multiply(w, x O) O {
	switch w := w.(type) {
	case B:
		return MultiplyBO(w, x)
	case F:
		return MultiplyFO(w, x)
	case I:
		return MultiplyIO(w, x)
	case S:
		return MultiplySO(w, x)
	case AB:
		return MultiplyABO(w, x)
	case AF:
		return MultiplyAFO(w, x)
	case AI:
		return MultiplyAIO(w, x)
	case AS:
		return MultiplyASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("×")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Multiply(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Multiply(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w && x
	case F:
		return B2F(w) * x
	case I:
		return B2I(w) * x
	case S:
		return strings.Repeat(x, B2I(w))
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) * x[i]
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], B2I(w))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MultiplyBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w * B2F(x)
	case F:
		return w * x
	case I:
		return w * F(x)
	case S:
		return strings.Repeat(x, I(math.Round(w)))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * F(x[i])
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], I(math.Round(w)))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MultiplyFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w * B2I(x)
	case F:
		return F(w) * x
	case I:
		return w * x
	case S:
		return strings.Repeat(x, w)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w * B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w * x[i]
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], w)
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MultiplyIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplySO(w S, x O) O {
	switch x := x.(type) {
	case B:
		return strings.Repeat(w, B2I(x))
	case F:
		return strings.Repeat(w, I(math.Round(x)))
	case I:
		return strings.Repeat(w, x)
	case AB:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w, B2I(x[i]))
		}
		return r
	case AF:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w, I(math.Round(x[i])))
		}
		return r
	case AI:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MultiplySO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] && x
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = B2F(w[i]) * x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) * x
		}
		return r
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(x, B2I(w[i]))
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w[i]) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) * x[i]
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], B2I(w[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MultiplyABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * B2F(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * x
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * F(x)
		}
		return r
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(x, I(math.Round(w[i])))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * F(x[i])
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], I(math.Round(w[i])))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MultiplyAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] * B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(w[i]) * x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] * x
		}
		return r
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(x, w[i])
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] * B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w[i]) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] * x[i]
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(x[i], w[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MultiplyAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func MultiplyASO(w AS, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(w[i], B2I(x))
		}
		return r
	case F:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(w[i], I(math.Round(x)))
		}
		return r
	case I:
		r := make(AS, len(w))
		for i := range r {
			r[i] = strings.Repeat(w[i], x)
		}
		return r
	case AB:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w[i], B2I(x[i]))
		}
		return r
	case AF:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w[i], I(math.Round(x[i])))
		}
		return r
	case AI:
		r := make(AS, len(x))
		for i := range r {
			r[i] = strings.Repeat(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MultiplyASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("×")
	}
}

func Divide(w, x O) O {
	switch w := w.(type) {
	case B:
		return DivideBO(w, x)
	case F:
		return DivideFO(w, x)
	case I:
		return DivideIO(w, x)
	case AB:
		return DivideABO(w, x)
	case AF:
		return DivideAFO(w, x)
	case AI:
		return DivideAIO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("÷")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Divide(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Divide(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return divide(B2F(w), B2F(x))
	case F:
		return divide(B2F(w), x)
	case I:
		return divide(B2F(w), F(x))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w), B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w), x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w), F(x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := DivideBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return divide(w, B2F(x))
	case F:
		return divide(w, x)
	case I:
		return divide(w, F(x))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w, B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w, x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w, F(x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := DivideFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return divide(F(w), B2F(x))
	case F:
		return divide(F(w), x)
	case I:
		return divide(F(w), F(x))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w), B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w), x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w), F(x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := DivideIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(B2F(w[i]), B2F(x))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(B2F(w[i]), x)
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(B2F(w[i]), F(x))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w[i]), B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w[i]), x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(B2F(w[i]), F(x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := DivideABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(w[i], B2F(x))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(w[i], x)
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(w[i], F(x))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w[i], B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w[i], x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(w[i], F(x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := DivideAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func DivideAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(F(w[i]), B2F(x))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(F(w[i]), x)
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = divide(F(w[i]), F(x))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w[i]), B2F(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w[i]), x[i])
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = divide(F(w[i]), F(x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := DivideAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("÷")
	}
}

func Minimum(w, x O) O {
	switch w := w.(type) {
	case B:
		return MinimumBO(w, x)
	case F:
		return MinimumFO(w, x)
	case I:
		return MinimumIO(w, x)
	case S:
		return MinimumSO(w, x)
	case AB:
		return MinimumABO(w, x)
	case AF:
		return MinimumAFO(w, x)
	case AI:
		return MinimumAIO(w, x)
	case AS:
		return MinimumASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("⌊")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Minimum(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Minimum(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w && x
	case F:
		return F(math.Min(float64(B2F(w)), float64(x)))
	case I:
		return minI(B2I(w), x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(B2F(w)), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(B2I(w), x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MinimumBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return F(math.Min(float64(w), float64(B2F(x))))
	case F:
		return F(math.Min(float64(w), float64(x)))
	case I:
		return F(math.Min(float64(w), float64(x)))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w), float64(B2F(x[i]))))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w), float64(x[i])))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w), float64(x[i])))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MinimumFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return minI(w, B2I(x))
	case F:
		return F(math.Min(float64(w), float64(x)))
	case I:
		return minI(w, x)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(w, B2I(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MinimumIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return minS(w, x)
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = minS(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MinimumSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] && x
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Min(float64(B2F(w[i])), float64(x)))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = minI(B2I(w[i]), x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(B2F(w[i])), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(B2I(w[i]), x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MinimumABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(B2F(x))))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x)))
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x)))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(B2F(x[i]))))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x[i])))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x[i])))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MinimumAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = minI(w[i], B2I(x))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x)))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = minI(w[i], x)
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(w[i], B2I(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Min(float64(w[i]), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = minI(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MinimumAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func MinimumASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = minS(w[i], x)
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = minS(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MinimumASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌊")
	}
}

func Maximum(w, x O) O {
	switch w := w.(type) {
	case B:
		return MaximumBO(w, x)
	case F:
		return MaximumFO(w, x)
	case I:
		return MaximumIO(w, x)
	case S:
		return MaximumSO(w, x)
	case AB:
		return MaximumABO(w, x)
	case AF:
		return MaximumAFO(w, x)
	case AI:
		return MaximumAIO(w, x)
	case AS:
		return MaximumASO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("⌈")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Maximum(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Maximum(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w || x
	case F:
		return F(math.Max(float64(B2F(w)), float64(x)))
	case I:
		return maxI(B2I(w), x)
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w || x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(B2F(w)), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(B2I(w), x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MaximumBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return F(math.Max(float64(w), float64(B2F(x))))
	case F:
		return F(math.Max(float64(w), float64(x)))
	case I:
		return F(math.Max(float64(w), float64(x)))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w), float64(B2F(x[i]))))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w), float64(x[i])))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w), float64(x[i])))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MaximumFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return maxI(w, B2I(x))
	case F:
		return F(math.Max(float64(w), float64(x)))
	case I:
		return maxI(w, x)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(w, B2I(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MaximumIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumSO(w S, x O) O {
	switch x := x.(type) {
	case S:
		return maxS(w, x)
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = maxS(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := MaximumSO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] || x
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Max(float64(B2F(w[i])), float64(x)))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = maxI(B2I(w[i]), x)
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] || x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(B2F(w[i])), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(B2I(w[i]), x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MaximumABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(B2F(x))))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x)))
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x)))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(B2F(x[i]))))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x[i])))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x[i])))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MaximumAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = maxI(w[i], B2I(x))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x)))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = maxI(w[i], x)
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(w[i], B2I(x[i]))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(math.Max(float64(w[i]), float64(x[i])))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = maxI(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MaximumAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func MaximumASO(w AS, x O) O {
	switch x := x.(type) {
	case S:
		r := make(AS, len(w))
		for i := range r {
			r[i] = maxS(w[i], x)
		}
		return r
	case AS:
		r := make(AS, len(x))
		for i := range r {
			r[i] = maxS(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := MaximumASO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("⌈")
	}
}

func And(w, x O) O {
	switch w := w.(type) {
	case B:
		return AndBO(w, x)
	case F:
		return AndFO(w, x)
	case I:
		return AndIO(w, x)
	case AB:
		return AndABO(w, x)
	case AF:
		return AndAFO(w, x)
	case AI:
		return AndAIO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("∧")
			}
			r := make(AO, len(w))
			for i := range r {
				v := And(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := And(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w && x
	case F:
		return B2F(w) * x
	case I:
		return B2I(w) * x
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w) * x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AndBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return w * B2F(x)
	case F:
		return w * x
	case I:
		return w * F(x)
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w * F(x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AndFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return w * B2I(x)
	case F:
		return F(w) * x
	case I:
		return w * x
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w * B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w * x[i]
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := AndIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] && x
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = B2F(w[i]) * x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = B2I(w[i]) * x
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] && x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = B2F(w[i]) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = B2I(w[i]) * x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AndABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * B2F(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * x
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = w[i] * F(x)
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * B2F(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * x[i]
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = w[i] * F(x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AndAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func AndAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] * B2I(x)
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = F(w[i]) * x
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = w[i] * x
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] * B2I(x[i])
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = F(w[i]) * x[i]
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = w[i] * x[i]
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := AndAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∧")
	}
}

func Or(w, x O) O {
	switch w := w.(type) {
	case B:
		return OrBO(w, x)
	case F:
		return OrFO(w, x)
	case I:
		return OrIO(w, x)
	case AB:
		return OrABO(w, x)
	case AF:
		return OrAFO(w, x)
	case AI:
		return OrAIO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("∨")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Or(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Or(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return w || x
	case F:
		return 1 - ((1 - B2F(w)) * (1 - x))
	case I:
		return 1 - ((1 - B2I(w)) * (1 - x))
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w || x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - B2F(w)) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - B2I(w)) * (1 - x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := OrBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return 1 - ((1 - w) * (1 - B2F(x)))
	case F:
		return 1 - ((1 - w) * (1 - x))
	case I:
		return 1 - ((1 - w) * F(1-x))
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w) * (1 - B2F(x[i])))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w) * F(1-x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := OrFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return 1 - ((1 - w) * (1 - B2I(x)))
	case F:
		return 1 - ((1 - F(w)) * (1 - x))
	case I:
		return 1 - ((1 - w) * (1 - x))
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w) * (1 - B2I(x[i])))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - F(w)) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w) * (1 - x[i]))
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := OrIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AB, len(w))
		for i := range r {
			r[i] = w[i] || x
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = 1 - ((1 - B2F(w[i])) * (1 - x))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = 1 - ((1 - B2I(w[i])) * (1 - x))
		}
		return r
	case AB:
		r := make(AB, len(x))
		for i := range r {
			r[i] = w[i] || x[i]
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - B2F(w[i])) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - B2I(w[i])) * (1 - x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := OrABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AF, len(w))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - B2F(x)))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - x))
		}
		return r
	case I:
		r := make(AF, len(w))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * F(1-x))
		}
		return r
	case AB:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - B2F(x[i])))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * F(1-x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := OrAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func OrAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - B2I(x)))
		}
		return r
	case F:
		r := make(AF, len(w))
		for i := range r {
			r[i] = 1 - ((1 - F(w[i])) * (1 - x))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - x))
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - B2I(x[i])))
		}
		return r
	case AF:
		r := make(AF, len(x))
		for i := range r {
			r[i] = 1 - ((1 - F(w[i])) * (1 - x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = 1 - ((1 - w[i]) * (1 - x[i]))
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := OrAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("∨")
	}
}

func Modulus(w, x O) O {
	switch w := w.(type) {
	case B:
		return ModulusBO(w, x)
	case F:
		return ModulusFO(w, x)
	case I:
		return ModulusIO(w, x)
	case AB:
		return ModulusABO(w, x)
	case AF:
		return ModulusAFO(w, x)
	case AI:
		return ModulusAIO(w, x)
	case AO:
		if isArray(x) {
			if length(x) != len(w) {
				return badlen("|")
			}
			r := make(AO, len(w))
			for i := range r {
				v := Modulus(w[i], at(x, i))
				e, ok := v.(E)
				if ok {
					return e
				}
				r[i] = v
			}
			return r
		}
		r := make(AO, len(w))
		for i := range r {
			v := Modulus(w[i], x)
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusBO(w B, x O) O {
	switch x := x.(type) {
	case B:
		return modulus(B2I(w), B2I(x))
	case F:
		return modulus(B2I(w), I(x))
	case I:
		return modulus(B2I(w), x)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w), B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w), I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w), x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := ModulusBO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusFO(w F, x O) O {
	switch x := x.(type) {
	case B:
		return modulus(I(w), B2I(x))
	case F:
		return modulus(I(w), I(x))
	case I:
		return modulus(I(w), x)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w), B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w), I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w), x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := ModulusFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusIO(w I, x O) O {
	switch x := x.(type) {
	case B:
		return modulus(w, B2I(x))
	case F:
		return modulus(w, I(x))
	case I:
		return modulus(w, x)
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w, B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w, I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w, x[i])
		}
		return r
	case AO:
		r := make([]O, len(x))
		for i := range r {
			v := ModulusIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusABO(w AB, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(B2I(w[i]), B2I(x))
		}
		return r
	case F:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(B2I(w[i]), I(x))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(B2I(w[i]), x)
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w[i]), B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w[i]), I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(B2I(w[i]), x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := ModulusABO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusAFO(w AF, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(I(w[i]), B2I(x))
		}
		return r
	case F:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(I(w[i]), I(x))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(I(w[i]), x)
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w[i]), B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w[i]), I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(I(w[i]), x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := ModulusAFO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}

func ModulusAIO(w AI, x O) O {
	switch x := x.(type) {
	case B:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(w[i], B2I(x))
		}
		return r
	case F:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(w[i], I(x))
		}
		return r
	case I:
		r := make(AI, len(w))
		for i := range r {
			r[i] = modulus(w[i], x)
		}
		return r
	case AB:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w[i], B2I(x[i]))
		}
		return r
	case AF:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w[i], I(x[i]))
		}
		return r
	case AI:
		r := make(AI, len(x))
		for i := range r {
			r[i] = modulus(w[i], x[i])
		}
		return r
	case AO:
		r := make(AO, len(x))
		for i := range r {
			v := ModulusAIO(w, x[i])
			e, ok := v.(E)
			if ok {
				return e
			}
			r[i] = v
		}
		return r
	case E:
		return w
	default:
		return badtype("|")
	}
}