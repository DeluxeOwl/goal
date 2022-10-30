// structural functions (Length, Reverse, Take, ...)

package main

import (
	"math"
	"sort"
)

// Length returns ≠x.
func Length(x V) I {
	switch x := x.(type) {
	case nil:
		return 0
	default:
		return I(x.Len())
	}
}

func reverse(x V) {
	switch x := x.(type) {
	case AB:
		for i := 0; i < len(x)/2; i++ {
			x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
		}
	case AF:
		for i := 0; i < len(x)/2; i++ {
			x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
		}
	case AI:
		for i := 0; i < len(x)/2; i++ {
			x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
		}
	case AS:
		for i := 0; i < len(x)/2; i++ {
			x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
		}
	case AV:
		for i := 0; i < len(x)/2; i++ {
			x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
		}
	case sort.Interface:
		sort.Reverse(x)
	}
}

// Reverse returns ⌽x.
func Reverse(x V) V {
	switch x := x.(type) {
	case Array:
		r := cloneShallow(x)
		reverse(r)
		return r
	default:
		return badtype("⌽")
	}
}

// Rotate returns w⌽x.
func Rotate(w, x V) V {
	i := 0
	switch w := w.(type) {
	case B:
		i = int(B2I(w))
	case I:
		i = int(w)
	case F:
		i = int(w)
	default:
		// TODO: improve error messages
		return badtype("w⌽")
	}
	lenx := int(Length(x))
	if lenx == 0 {
		return x
	}
	i %= lenx
	if i < 0 {
		i += lenx
	}
	switch x := x.(type) {
	case AB:
		r := make(AB, lenx)
		for j := 0; j < lenx; j++ {
			r[j] = x[(j+i)%lenx]
		}
		return r
	case AF:
		r := make(AF, lenx)
		for j := 0; j < lenx; j++ {
			r[j] = x[(j+i)%lenx]
		}
		return r
	case AI:
		r := make(AI, lenx)
		for j := 0; j < lenx; j++ {
			r[j] = x[(j+i)%lenx]
		}
		return r
	case AS:
		r := make(AS, lenx)
		for j := 0; j < lenx; j++ {
			r[j] = x[(j+i)%lenx]
		}
		return r
	case AV:
		r := make(AV, lenx)
		for j := 0; j < lenx; j++ {
			r[j] = x[(j+i)%lenx]
		}
		return r
	default:
		return badtype("⌽x")
	}
}

// First returns ↑x.
func First(x V) V {
	switch x := x.(type) {
	case Array:
		if x.Len() == 0 {
			switch x.(type) {
			case AB:
				return B(false)
			case AF:
				return F(0)
			case AI:
				return I(0)
			case AS:
				return S("")
			default:
				return V(nil)
			}
		}
		return x.At(0)
	default:
		return x
	}
}

// Tail returns ↓x.
func Tail(x V) V {
	x = toArray(x)
	switch x := x.(type) {
	case Array:
		if x.Len() == 0 {
			return badlen("↓")
		}
		return x.Slice(1, x.Len())
	default:
		return badtype("↓")
	}
}

// Drop returns i_x.
func Drop(w, x V) V {
	i := 0
	switch w := w.(type) {
	case B:
		i = int(B2I(w))
	case I:
		i = int(w)
	case F:
		i = int(math.Round(float64(w)))
	default:
		// TODO: improve error messages
		return badtype("w↓")
	}
	x = toArray(x)
	switch x := x.(type) {
	case Array:
		switch {
		case i >= 0:
			if i > x.Len() {
				i = x.Len()
			}
			return x.Slice(i, x.Len())
		default:
			i = x.Len() + i
			if i < 0 {
				i = 0
			}
			return x.Slice(0, i)
		}
	default:
		return x
	}
}

// Take returns i#x.
func Take(w, x V) V {
	i := 0
	switch w := w.(type) {
	case B:
		i = int(B2I(w))
	case I:
		i = int(w)
	case F:
		i = int(math.Round(float64(w)))
	default:
		// TODO: improve error messages
		return badtype("w↑")
	}
	x = toArray(x)
	switch x := x.(type) {
	case Array:
		switch {
		case i >= 0:
			if i > x.Len() {
				return growArray(x, i)
			}
			return x.Slice(0, i)
		default:
			if i < -x.Len() {
				return growArray(x, i)
			}
			return x.Slice(x.Len()+i, x.Len())
		}
	default:
		return x
	}
}

// ShiftBefore returns w»x.
func ShiftBefore(w, x V) V {
	w = toArray(w)
	max := int(minI(Length(w), Length(x)))
	if max == 0 {
		return x
	}
	switch x := x.(type) {
	case AB:
		switch w := w.(type) {
		case AB:
			r := make(AB, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			copy(r[max:], x[:len(x)-max])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i] = float64(B2F(B(x[i-max])))
			}
			return r
		case AI:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i] = int(B2I(B(x[i-max])))
			}
			return r
		default:
			return badtype("» : type mismatch")
		}
	case AF:
		switch w := w.(type) {
		case AB:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[i] = float64(B2F(B(w[i])))
			}
			copy(r[max:], x[:len(x)-max])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			copy(r[max:], x[:len(x)-max])
			return r
		case AI:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[i] = float64(w[i])
			}
			copy(r[max:], x[:len(x)-max])
			return r
		default:
			return badtype("» : type mismatch")
		}
	case AI:
		switch w := w.(type) {
		case AB:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[i] = int(B2I(B(w[i])))
			}
			copy(r[max:], x[:len(x)-max])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i] = float64(x[i-max])
			}
			return r
		case AI:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			copy(r[max:], x[:len(x)-max])
			return r
		default:
			return badtype("» : type mismatch")
		}
	case AS:
		switch w := w.(type) {
		case AS:
			r := make(AS, len(x))
			for i := 0; i < max; i++ {
				r[i] = w[i]
			}
			copy(r[max:], x[:len(x)-max])
			return r
		default:
			return badtype("» : type mismatch")
		}
	case AV:
		switch w := w.(type) {
		case Array:
			r := make(AV, len(x))
			for i := 0; i < max; i++ {
				r[i] = w.At(i)
			}
			copy(r[max:], x[:len(x)-max])
			return r
		default:
			return badtype("» : type mismatch")
		}
	default:
		return badtype("» : x must be an array")
	}
}

// Nudge returns »x.
func Nudge(x V) V {
	switch x := x.(type) {
	case AB:
		r := make(AB, len(x))
		copy(r[1:], x[0:len(x)-1])
		return r
	case AI:
		r := make(AI, len(x))
		copy(r[1:], x[0:len(x)-1])
		return r
	case AF:
		r := make(AF, len(x))
		copy(r[1:], x[0:len(x)-1])
		return r
	case AS:
		r := make(AS, len(x))
		copy(r[1:], x[0:len(x)-1])
		return r
	case AV:
		r := make(AV, len(x))
		copy(r[1:], x[0:len(x)-1])
		return r
	default:
		return badtype("» : x must be an array")
	}
}

// ShiftAfter returns w«x.
func ShiftAfter(w, x V) V {
	w = toArray(w)
	max := int(minI(Length(w), Length(x)))
	if max == 0 {
		return x
	}
	switch x := x.(type) {
	case AB:
		switch w := w.(type) {
		case AB:
			r := make(AB, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			copy(r[:len(x)-max], x[max:])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i-max] = float64(B2F(B(x[i])))
			}
			return r
		case AI:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i-max] = int(B2I(B(x[i])))
			}
			return r
		default:
			return badtype("« : type mismatch")
		}
	case AF:
		switch w := w.(type) {
		case AB:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = float64(B2F(B(w[i])))
			}
			copy(r[:len(x)-max], x[max:])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			copy(r[:len(x)-max], x[max:])
			return r
		case AI:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = float64(w[i])
			}
			copy(r[:len(x)-max], x[max:])
			return r
		default:
			return badtype("« : type mismatch")
		}
	case AI:
		switch w := w.(type) {
		case AB:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = int(B2I(B(w[i])))
			}
			copy(r[:len(x)-max], x[max:])
			return r
		case AF:
			r := make(AF, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			for i := max; i < len(x); i++ {
				r[i-max] = float64(x[max])
			}
			return r
		case AI:
			r := make(AI, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			copy(r[:len(x)-max], x[max:])
			return r
		default:
			return badtype("« : type mismatch")
		}
	case AS:
		switch w := w.(type) {
		case AS:
			r := make(AS, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w[i]
			}
			copy(r[:len(x)-max], x[max:])
			return r
		default:
			return badtype("« : type mismatch")
		}
	case AV:
		switch w := w.(type) {
		case Array:
			r := make(AV, len(x))
			for i := 0; i < max; i++ {
				r[len(x)-1-i] = w.At(i)
			}
			copy(r[:len(x)-max], x[max:])
			return r
		default:
			return badtype("« : type mismatch")
		}
	default:
		return badtype("« : x must be an array")
	}
}

// NudgeBack returns «x.
func NudgeBack(x V) V {
	if Length(x) == 0 {
		return x
	}
	switch x := x.(type) {
	case AB:
		r := make(AB, len(x))
		copy(r[0:len(x)-1], x[1:])
		return r
	case AI:
		r := make(AI, len(x))
		copy(r[0:len(x)-1], x[1:])
		return r
	case AF:
		r := make(AF, len(x))
		copy(r[0:len(x)-1], x[1:])
		return r
	case AS:
		r := make(AS, len(x))
		copy(r[0:len(x)-1], x[1:])
		return r
	case AV:
		r := make(AV, len(x))
		copy(r[0:len(x)-1], x[1:])
		return r
	default:
		return badtype("« : x must be an array")
	}
}

// Flip returns +x.
func Flip(x V) V {
	x = toArray(x)
	x = canonical(x) // XXX really?
	switch x := x.(type) {
	case AV:
		cols := len(x)
		if cols == 0 {
			// (+⟨⟩) ≡ ⋈⟨⟩
			return AV{x}
		}
		lines := -1
		for _, o := range x {
			nl := int(Length(o))
			if !isArray(o) {
				continue
			}
			switch {
			case lines < 0:
				lines = nl
			case nl >= 1 && nl != lines:
				return badlen("+")
			}
		}
		t := aType(x)
		switch {
		case lines <= 0:
			// (+⟨⟨⟩,…⟩) ≡ ⟨⟩
			// TODO: error if atoms?
			return x[0]
		case lines == 1:
			switch t {
			case tB, tAB:
				return AV{flipAB(x)}
			case tF, tAF:
				return AV{flipAF(x)}
			case tI, tAI:
				return AV{flipAI(x)}
			case tS, tAS:
				return AV{flipAS(x)}
			default:
				return AV{flipAO(x)}
			}
		default:
			switch t {
			case tB, tAB:
				return flipAOAB(x, lines)
			case tF, tAF:
				return flipAOAF(x, lines)
			case tI, tAI:
				return flipAOAI(x, lines)
			case tS, tAS:
				return flipAOAS(x, lines)
			default:
				return flipAOAO(x, lines)
			}
		}
	default:
		return AV{x}
	}
}

func flipAB(x AV) AB {
	r := make(AB, len(x))
	for i, y := range x {
		switch y := y.(type) {
		case B:
			r[i] = bool(y)
		case AB:
			r[i] = y[0]
		}
	}
	return r
}

func flipAOAB(x AV, lines int) AV {
	r := make(AV, lines)
	a := make(AB, lines*len(x))
	for j := range r {
		q := a[j*len(x) : (j+1)*len(x)]
		for i, y := range x {
			switch y := y.(type) {
			case B:
				q[i] = bool(y)
			case AB:
				q[i] = y[j]
			}
		}
		r[j] = q
	}
	return r
}

func flipAF(x AV) AF {
	r := make(AF, len(x))
	for i, y := range x {
		switch y := y.(type) {
		case B:
			r[i] = float64(B2F(y))
		case AB:
			r[i] = float64(B2F(B(y[0])))
		case F:
			r[i] = float64(y)
		case AF:
			r[i] = y[0]
		case I:
			r[i] = float64(y)
		case AI:
			r[i] = float64(y[0])
		}
	}
	return r
}

func flipAOAF(x AV, lines int) AV {
	r := make(AV, lines)
	a := make(AF, lines*len(x))
	for j := range r {
		q := a[j*len(x) : (j+1)*len(x)]
		for i, y := range x {
			switch y := y.(type) {
			case B:
				q[i] = float64(B2F(y))
			case AB:
				q[i] = float64(B2F(B(y[j])))
			case F:
				q[i] = float64(y)
			case AF:
				q[i] = y[j]
			case I:
				q[i] = float64(y)
			case AI:
				q[i] = float64(y[j])
			}
		}
		r[j] = q
	}
	return r
}

func flipAI(x AV) AI {
	r := make(AI, len(x))
	for i, y := range x {
		switch y := y.(type) {
		case B:
			r[i] = int(B2I(y))
		case AB:
			r[i] = int(B2I(B(y[0])))
		case I:
			r[i] = int(y)
		case AI:
			r[i] = y[0]
		}
	}
	return r
}

func flipAOAI(x AV, lines int) AV {
	r := make(AV, lines)
	a := make(AI, lines*len(x))
	for j := range r {
		q := a[j*len(x) : (j+1)*len(x)]
		for i, y := range x {
			switch y := y.(type) {
			case B:
				q[i] = int(B2I(y))
			case AB:
				q[i] = int(B2I(B(y[j])))
			case I:
				q[i] = int(y)
			case AI:
				q[i] = y[j]
			}
		}
		r[j] = q
	}
	return r
}

func flipAS(x AV) AS {
	r := make(AS, len(x))
	for i, y := range x {
		switch y := y.(type) {
		case S:
			r[i] = string(y)
		case AS:
			r[i] = y[0]
		}
	}
	return r
}

func flipAOAS(x AV, lines int) AV {
	r := make(AV, lines)
	a := make(AS, lines*len(x))
	for j := range r {
		q := a[j*len(x) : (j+1)*len(x)]
		for i, y := range x {
			switch y := y.(type) {
			case S:
				q[i] = string(y)
			case AS:
				q[i] = y[j]
			}
		}
		r[j] = q
	}
	return r
}

func flipAO(x AV) AV {
	r := make(AV, len(x))
	for i, y := range x {
		switch y := y.(type) {
		case Array:
			r[i] = y.At(0)
		default:
			r[i] = y
		}
	}
	return r
}

func flipAOAO(x AV, lines int) AV {
	r := make(AV, lines)
	a := make(AV, lines*len(x))
	for j := range r {
		q := a[j*len(x) : (j+1)*len(x)]
		for i, y := range x {
			switch y := y.(type) {
			case Array:
				q[i] = y.At(j)
			default:
				q[i] = y
			}
		}
		r[j] = q
	}
	return r
}

// JoinTo returns w,x.
func JoinTo(w, x V) V {
	switch w := w.(type) {
	case B:
		return joinToB(w, x, true)
	case F:
		return joinToF(w, x, true)
	case I:
		return joinToI(w, x, true)
	case S:
		return joinToS(w, x, true)
	case AB:
		return joinToAB(x, w, false)
	case AF:
		return joinToAF(x, w, false)
	case AI:
		return joinToAI(x, w, false)
	case AS:
		return joinToAS(x, w, false)
	case AV:
		return joinToAO(x, w, false)
	default:
		switch x := x.(type) {
		case Array:
			return joinAtomToArray(w, x, true)
		default:
			return AV{w, x}
		}
	}
}

func joinToB(w B, x V, left bool) V {
	switch x := x.(type) {
	case B:
		if left {
			return AB{bool(w), bool(x)}
		}
		return AB{bool(x), bool(w)}
	case F:
		if left {
			return AF{float64(B2F(w)), float64(x)}
		}
		return AF{float64(x), float64(B2F(w))}
	case I:
		if left {
			return AI{int(B2I(w)), int(x)}
		}
		return AI{int(x), int(B2I(w))}
	case S:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case AB:
		return joinToAB(w, x, left)
	case AF:
		return joinToAF(w, x, left)
	case AI:
		return joinToAI(w, x, left)
	case AS:
		return joinToAS(w, x, left)
	case AV:
		return joinToAO(w, x, left)
	default:
		return AV{w, x}
	}
}

func joinToI(w I, x V, left bool) V {
	switch x := x.(type) {
	case B:
		if left {
			return AI{int(w), int(B2I(x))}
		}
		return AI{int(B2I(x)), int(w)}
	case F:
		if left {
			return AF{float64(w), float64(x)}
		}
		return AF{float64(x), float64(w)}
	case I:
		if left {
			return AI{int(w), int(x)}
		}
		return AI{int(x), int(w)}
	case S:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case AB:
		return joinToAB(w, x, left)
	case AF:
		return joinToAF(w, x, left)
	case AI:
		return joinToAI(w, x, left)
	case AS:
		return joinToAS(w, x, left)
	case AV:
		return joinToAO(w, x, left)
	default:
		return AV{w, x}
	}
}

func joinToF(w F, x V, left bool) V {
	switch x := x.(type) {
	case B:
		if left {
			return AF{float64(w), float64(B2F(x))}
		}
		return AF{float64(B2F(x)), float64(w)}
	case F:
		if left {
			return AF{float64(w), float64(x)}
		}
		return AF{float64(x), float64(w)}
	case I:
		if left {
			return AF{float64(w), float64(x)}
		}
		return AF{float64(x), float64(w)}
	case S:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case AB:
		return joinToAB(w, x, left)
	case AF:
		return joinToAF(w, x, left)
	case AI:
		return joinToAI(w, x, left)
	case AS:
		return joinToAS(w, x, left)
	case AV:
		return joinToAO(w, x, left)
	default:
		return AV{w, x}
	}
}

func joinToS(w S, x V, left bool) V {
	switch x := x.(type) {
	case B:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case F:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case I:
		if left {
			return AV{w, x}
		}
		return AV{x, w}
	case S:
		if left {
			return AS{string(w), string(x)}
		}
		return AS{string(x), string(w)}
	case AB:
		return joinToAB(w, x, left)
	case AF:
		return joinToAF(w, x, left)
	case AI:
		return joinToAI(w, x, left)
	case AS:
		return joinToAS(w, x, left)
	case AV:
		return joinToAO(w, x, left)
	default:
		return AV{w, x}
	}
}

func joinToAO(w V, x AV, left bool) V {
	switch w := w.(type) {
	case Array:
		if left {
			return joinArrays(w, x)
		}
		return joinArrays(x, w)
	default:
		r := make(AV, len(x)+1)
		if left {
			r[0] = w
			copy(r[1:], x)
		} else {
			r[len(r)-1] = w
			copy(r[:len(r)-1], x)
		}
		return r
	}
}

func joinArrays(w, x Array) AV {
	r := make(AV, x.Len()+w.Len())
	for i := 0; i < w.Len(); i++ {
		r[i] = w.At(i)
	}
	for i := w.Len(); i < len(r); i++ {
		r[i] = x.At(i - w.Len())
	}
	return r
}

func joinAtomToArray(w V, x Array, left bool) AV {
	r := make(AV, x.Len()+1)
	if left {
		r[0] = w
		for i := 1; i < len(r); i++ {
			r[i] = x.At(i - 1)
		}
	} else {
		r[len(r)-1] = w
		for i := 0; i < len(r)-1; i++ {
			r[i] = x.At(i)
		}
	}
	return r
}

func joinToAS(w V, x AS, left bool) V {
	switch w := w.(type) {
	case S:
		r := make(AS, len(x)+1)
		if left {
			r[0] = string(w)
			copy(r[1:], x)
		} else {
			r[len(r)-1] = string(w)
			copy(r[:len(r)-1], x)
		}
		return r
	case AS:
		r := make(AS, len(x)+len(w))
		if left {
			copy(r[:len(w)], w)
			copy(r[len(w):], x)
		} else {
			copy(r[:len(x)], x)
			copy(r[len(x):], w)
		}
		return r
	case Array:
		if left {
			return joinArrays(w, x)
		}
		return joinArrays(x, w)
	default:
		return joinAtomToArray(w, x, left)
	}
}

func joinToAB(w V, x AB, left bool) V {
	switch w := w.(type) {
	case B:
		r := make(AB, len(x)+1)
		if left {
			r[0] = bool(w)
			copy(r[1:], x)
		} else {
			r[len(r)-1] = bool(w)
			copy(r[:len(r)-1], x)
		}
		return r
	case F:
		r := make(AF, len(x)+1)
		if left {
			r[0] = float64(w)
			for i := 1; i < len(r); i++ {
				r[i] = float64(B2F(B(x[i-1])))
			}
		} else {
			r[len(r)-1] = float64(w)
			for i := 0; i < len(r); i++ {
				r[i] = float64(B2F(B(x[i])))
			}
		}
		return r
	case I:
		r := make(AI, len(x)+1)
		if left {
			r[0] = int(w)
			for i := 1; i < len(r); i++ {
				r[i] = int(B2I(B(x[i-1])))
			}
		} else {
			r[len(r)-1] = int(w)
			for i := 0; i < len(r); i++ {
				r[i] = int(B2I(B(x[i])))
			}
		}
		return r
	case AB:
		if left {
			return joinABAB(w, x)
		}
		return joinABAB(x, w)
	case AI:
		if left {
			return joinAIAB(w, x)
		}
		return joinABAI(x, w)
	case AF:
		if left {
			return joinAFAB(w, x)
		}
		return joinABAF(x, w)
	case Array:
		if left {
			return joinArrays(w, x)
		}
		return joinArrays(x, w)
	default:
		return joinAtomToArray(w, x, left)
	}
}

func joinToAI(w V, x AI, left bool) V {
	switch w := w.(type) {
	case B:
		r := make(AI, len(x)+1)
		if left {
			r[0] = int(B2I(w))
			copy(r[1:], x)
		} else {
			r[len(r)-1] = int(B2I(w))
			copy(r[:len(r)-1], x)
		}
		return r
	case F:
		r := make(AF, len(x)+1)
		if left {
			r[0] = float64(w)
			for i := 1; i < len(r); i++ {
				r[i] = float64(x[i-1])
			}
		} else {
			r[len(r)-1] = float64(w)
			for i := 0; i < len(r)-1; i++ {
				r[i] = float64(x[i])
			}
		}
		return r
	case I:
		r := make(AI, len(x)+1)
		if left {
			r[0] = int(w)
			copy(r[1:], x)
		} else {
			r[len(r)-1] = int(w)
			copy(r[:len(r)-1], x)
		}
		return r
	case AB:
		if left {
			return joinABAI(w, x)
		}
		return joinAIAB(x, w)
	case AI:
		if left {
			return joinAIAI(w, x)
		}
		return joinAIAI(x, w)
	case AF:
		if left {
			return joinAFAI(w, x)
		}
		return joinAIAF(x, w)
	case Array:
		if left {
			return joinArrays(w, x)
		}
		return joinArrays(x, w)
	default:
		return joinAtomToArray(w, x, left)
	}
}

func joinToAF(w V, x AF, left bool) V {
	switch w := w.(type) {
	case B:
		r := make(AF, len(x)+1)
		if left {
			r[0] = float64(B2F(w))
			copy(r[1:], x)
		} else {
			r[len(r)-1] = float64(B2F(w))
			copy(r[:len(r)-1], x)
		}
		return r
	case F:
		r := make(AF, len(x)+1)
		if left {
			r[0] = float64(w)
			copy(r[1:], x)
		} else {
			r[len(r)-1] = float64(w)
			copy(r[:len(r)-1], x)
		}
		return r
	case I:
		r := make(AF, len(x)+1)
		if left {
			r[0] = float64(w)
			copy(r[1:], x)
		} else {
			r[len(r)-1] = float64(w)
			copy(r[:len(r)-1], x)
		}
		return r
	case AB:
		if left {
			return joinABAF(w, x)
		}
		return joinAFAB(x, w)
	case AI:
		if left {
			return joinAIAF(w, x)
		}
		return joinAFAI(x, w)
	case AF:
		if left {
			return joinAFAF(w, x)
		}
		return joinAFAF(x, w)
	case Array:
		if left {
			return joinArrays(w, x)
		}
		return joinArrays(x, w)
	default:
		return joinAtomToArray(w, x, left)
	}
}

func joinABAB(w AB, x AB) AB {
	r := make(AB, len(x)+len(w))
	copy(r[:len(w)], w)
	copy(r[len(w):], x)
	return r
}

func joinAIAI(w AI, x AI) AI {
	r := make(AI, len(x)+len(w))
	copy(r[:len(w)], w)
	copy(r[len(w):], x)
	return r
}

func joinAFAF(w AF, x AF) AF {
	r := make(AF, len(x)+len(w))
	copy(r[:len(w)], w)
	copy(r[len(w):], x)
	return r
}

func joinABAI(w AB, x AI) AI {
	r := make(AI, len(w)+len(x))
	for i := 0; i < len(w); i++ {
		r[i] = int(B2I(B(w[i])))
	}
	copy(r[len(w):], x)
	return r
}

func joinAIAB(w AI, x AB) AI {
	r := make(AI, len(w)+len(x))
	copy(r[:len(w)], w)
	for i := len(w); i < len(r); i++ {
		r[i] = int(B2I(B(x[i-len(w)])))
	}
	return r
}

func joinABAF(w AB, x AF) AF {
	r := make(AF, len(w)+len(x))
	for i := 0; i < len(w); i++ {
		r[i] = float64(B2F(B(w[i])))
	}
	copy(r[len(w):], x)
	return r
}

func joinAFAB(w AF, x AB) AF {
	r := make(AF, len(w)+len(x))
	copy(r[:len(w)], w)
	for i := len(w); i < len(r); i++ {
		r[i] = float64(B2F(B(x[i-len(w)])))
	}
	return r
}

func joinAIAF(w AI, x AF) AF {
	r := make(AF, len(w)+len(x))
	for i := 0; i < len(w); i++ {
		r[i] = float64(w[i])
	}
	copy(r[len(w):], x)
	return r
}

func joinAFAI(w AF, x AI) AF {
	r := make(AF, len(w)+len(x))
	copy(r[:len(w)], w)
	for i := len(w); i < len(r); i++ {
		r[i] = float64(x[i-len(w)])
	}
	return r
}

// Enlist returns ,x.
func Enlist(x V) V {
	switch x := x.(type) {
	case B:
		return AB{bool(x)}
	case F:
		return AF{float64(x)}
	case I:
		return AI{int(x)}
	case S:
		return AS{string(x)}
	default:
		return AV{x}
	}
}

//// Pair returns w⋈x.
//func Pair(w, x O) O {
//switch w := w.(type) {
//case B:
//switch x := x.(type) {
//case B:
//return AB{bool(w), bool(x)}
//case I:
//return AI{B2I(w), x}
//case F:
//return AF{B2F(w), x}
//default:
//return AO{w, x}
//}
//case F:
//switch x := x.(type) {
//case B:
//return AF{w, B2F(x)}
//case I:
//return AF{w, F(x)}
//case F:
//return AF{w, x}
//default:
//return AO{w, x}
//}
//case I:
//switch x := x.(type) {
//case B:
//return AI{w, B2I(x)}
//case I:
//return AI{w, x}
//case F:
//return AF{F(w), x}
//default:
//return AO{w, x}
//}
//case S:
//switch x := x.(type) {
//case S:
//return AS{w, x}
//default:
//return AO{w, x}
//}
//default:
//return AO{w, x}
//}
//}

// Windows returns w↕x.
func Windows(w, x V) V {
	i := 0
	switch w := w.(type) {
	case B:
		i = int(B2I(w))
	case I:
		i = int(w)
	case F:
		i = int(math.Round(float64(w)))
	default:
		// TODO: improve error messages
		return badtype("↕ : w must be un integer")
	}
	switch x := x.(type) {
	case Array:
		if i <= 0 || i >= x.Len()+1 {
			return badtype("↕ : w must be between 0 and 1+≠x")
		}
		r := make(AV, 1+x.Len()-i)
		for j := range r {
			r[j] = x.Slice(j, j+i)
		}
		return r
	default:
		return badtype("↕ : x must be an array")
	}
}

// Group returns ⊔x.
func Group(x V) V {
	if Length(x) == 0 {
		return AV{}
	}
	// TODO: optimize allocations
	switch x := x.(type) {
	case AB:
		_, max := minMaxB(x)
		r := make(AV, B2I(max)+1)
		for i := range r {
			r[i] = AI{}
		}
		for i, v := range x {
			j := B2I(B(v))
			rj := r[j].(AI)
			r[j] = append(rj, i)
		}
		return r
	case AI:
		min, max := minMax(x)
		if min < 0 {
			return badtype("⊔ : x must not contain negative values")
		}
		r := make(AV, max+1)
		for i := range r {
			r[i] = AI{}
		}
		for i, j := range x {
			rj := r[j].(AI)
			r[j] = append(rj, i)
		}
		return r
		// TODO: AF and AO
	default:
		return badtype("⊔ : x must be a non negative integer array")
	}
}
