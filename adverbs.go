package goal

import (
	//"fmt"
	"strings"
)

func fold2(ctx *Context, args []V) V {
	f := args[1]
	switch f.Kind {
	case Variadic:
		switch f.variadic() {
		case vAdd:
			return fold2vAdd(args[0])
		}
	}
	if !f.IsFunction() {
		if f.IsI() {
			return fold2Decode(f, args[0])
		}
		if f.IsF() {
			return fold2Decode(f, args[0])
		}
		switch fv := f.Value.(type) {
		case S:
			return fold2Join(fv, args[0])
		case *AB, *AI, *AF:
			return fold2Decode(f, args[0])
		default:
			return panicType("F/x", "F", f)
		}
	}
	if f.Rank(ctx) != 2 {
		// TODO: converge
		return panicf("F/x : F rank is %d (expected 2)", f.Rank(ctx))
	}
	x := args[0]
	switch xv := x.Value.(type) {
	case array:
		if xv.Len() == 0 {
			f, ok := f.Value.(zeroFun)
			if ok {
				return f.zero()
			}
			return NewI(0)
		}
		r := xv.at(0)
		for i := 1; i < xv.Len(); i++ {
			ctx.push(xv.at(i))
			ctx.push(r)
			r = ctx.applyN(f, 2)
		}
		return r
	default:
		return x
	}
}

func fold2vAdd(x V) V {
	switch xv := x.Value.(type) {
	case *AB:
		n := int64(0)
		for _, b := range xv.Slice {
			if b {
				n++
			}
		}
		return NewI(n)
	case *AI:
		n := int64(0)
		for _, xi := range xv.Slice {
			n += xi
		}
		return NewI(n)
	case *AF:
		n := 0.0
		for _, xi := range xv.Slice {
			n += xi
		}
		return NewF(n)
	case *AS:
		if xv.Len() == 0 {
			return NewS("")
		}
		n := 0
		for _, s := range xv.Slice {
			n += len(s)
		}
		var b strings.Builder
		b.Grow(n)
		for _, s := range xv.Slice {
			b.WriteString(s)
		}
		return NewS(b.String())
	case *AV:
		if xv.Len() == 0 {
			return NewI(0)
		}
		r := xv.At(0)
		for _, xi := range xv.Slice[1:] {
			r = add(r, xi)
		}
		return r
	default:
		return x
	}
}

func fold2Join(sep S, x V) V {
	switch xv := x.Value.(type) {
	case S:
		return x
	case *AS:
		return NewS(strings.Join([]string(xv.Slice), string(sep)))
	case *AV:
		//assertCanonical(xv)
		return panicf("s/x : x not a string array (%s)", x.Type())
	default:
		return panicf("s/x : x not a string array (%s)", x.Type())
	}
}

func fold2Decode(f V, x V) V {
	if f.IsI() {
		if x.IsI() {
			return x
		}
		if x.IsF() {
			if !isI(x.F()) {
				return panicf("i/x : x non-integer (%g)", x.F())
			}
			return NewI(int64(x.F()))
		}
		switch xv := x.Value.(type) {
		case *AI:
			var r, n int64 = 0, 1
			for i := xv.Len() - 1; i >= 0; i-- {
				r += xv.At(i) * n
				n *= f.I()
			}
			return NewI(r)
		case *AB:
			var r, n int64 = 0, 1
			for i := xv.Len() - 1; i >= 0; i-- {
				r += B2I(xv.At(i)) * n
				n *= f.I()
			}
			return NewI(r)
		case *AF:
			aix := toAI(xv)
			if aix.isPanic() {
				return aix
			}
			return fold2Decode(f, aix)
		case *AV:
			r := make([]V, xv.Len())
			for i, xi := range xv.Slice {
				r[i] = fold2Decode(f, xi)
				if r[i].isPanic() {
					return r[i]
				}
			}
			return canonicalV(NewAV(r))
		default:
			return panicType("i/x", "x", x)
		}

	}
	if f.IsF() {
		if !isI(f.F()) {
			return panicf("i/x : i non-integer (%g)", f.F())
		}
		return fold2Decode(NewI(int64(f.F())), x)
	}
	switch fv := f.Value.(type) {
	case *AB:
		return fold2Decode(fromABtoAI(fv), x)
	case *AI:
		if x.IsI() {
			var r, n int64 = 0, 1
			for i := fv.Len() - 1; i >= 0; i-- {
				r += x.I() * n
				n *= fv.At(i)
			}
			return NewI(r)
		}
		if x.IsF() {
			if !isI(x.F()) {
				return panicf("I/x : x non-integer (%g)", x.F())
			}
			return fold2Decode(f, NewI(int64(x.F())))
		}
		switch xv := x.Value.(type) {
		case *AI:
			if fv.Len() != xv.Len() {
				return panicf("I/x : length mismatch: %d (#I) %d (#x)", fv.Len(), xv.Len())
			}
			var r, n int64 = 0, 1
			for i := xv.Len() - 1; i >= 0; i-- {
				r += xv.At(i) * n
				n *= fv.At(i)
			}
			return NewI(r)
		case *AB:
			return fold2Decode(f, fromABtoAI(xv))
		case *AF:
			aix := toAI(xv)
			if aix.isPanic() {
				return aix
			}
			return fold2Decode(f, aix)
		case *AV:
			r := make([]V, xv.Len())
			for i, xi := range xv.Slice {
				r[i] = fold2Decode(f, xi)
				if r[i].isPanic() {
					return r[i]
				}
			}
			return canonicalV(NewAV(r))
		default:
			return panicType("I/x", "x", x)
		}
	case *AF:
		aif := toAI(fv)
		if aif.isPanic() {
			return aif
		}
		return fold2Decode(aif, x)
	default:
		// should not happen
		return panicType("I/x", "I", f)
	}
}

func fold3(ctx *Context, args []V) V {
	f := args[1]
	if !f.IsFunction() {
		return panicf("x F/y : F not a function (%s)", f.Type())
	}
	if f.Rank(ctx) != 2 {
		return fold3While(ctx, args)
	}
	y := args[0]
	switch yv := y.Value.(type) {
	case array:
		r := args[2]
		if yv.Len() == 0 {
			return r
		}
		for i := 0; i < yv.Len(); i++ {
			ctx.push(yv.at(i))
			ctx.push(r)
			r = ctx.applyN(f, 2)
			if r.isPanic() {
				return r
			}
		}
		return canonicalV(r)
	default:
		ctx.push(y)
		ctx.push(args[2])
		return ctx.applyN(f, 2)
	}
}

func fold3While(ctx *Context, args []V) V {
	f := args[1]
	x := args[2]
	y := args[0]
	if x.IsI() {
		return fold3doTimes(ctx, x.I(), f, y)
	}
	if x.IsF() {
		if !isI(x.F()) {
			return panicf("n f/y : non-integer n (%g)", x.F())
		}
		return fold3doTimes(ctx, int64(x.F()), f, y)
	}
	if x.IsFunction() {
		for {
			ctx.push(y)
			y.rcincr()
			cond := ctx.applyN(x, 1)
			y.rcdecr()
			if cond.isPanic() {
				return cond
			}
			if !isTrue(cond) {
				return y
			}
			ctx.push(y)
			y = ctx.applyN(f, 1)
			if y.isPanic() {
				return y
			}
		}
	}
	return panicType("x f/y", "x", x)
}

func fold3doTimes(ctx *Context, n int64, f, y V) V {
	for i := int64(0); i < n; i++ {
		ctx.push(y)
		y = ctx.applyN(f, 1)
		if y.isPanic() {
			return y
		}
	}
	return y
}

func scan2(ctx *Context, f, x V) V {
	if !f.IsFunction() {
		if f.IsI() {
			return scan2Encode(f, x)
		}
		if f.IsF() {
			return scan2Encode(f, x)
		}
		switch fv := f.Value.(type) {
		case S:
			return scan2Split(fv, x)
		case *AB, *AI, *AF:
			return scan2Encode(f, x)
		default:
			return panicType("f\\x", "f", f)
		}
	}
	if f.Rank(ctx) != 2 {
		// TODO: converge
		return panicf("f\\x : f rank is %d (expected 2)", f.Rank(ctx))
	}
	switch xv := x.Value.(type) {
	case array:
		if xv.Len() == 0 {
			ff, ok := f.Value.(zeroFun)
			if ok {
				return ff.zero()
			}
			return NewI(0)
		}
		r := []V{xv.at(0)}
		for i := 1; i < xv.Len(); i++ {
			ctx.push(xv.at(i))
			last := r[len(r)-1]
			ctx.push(last)
			last.rcincr()
			next := ctx.applyN(f, 2)
			last.rcdecr()
			if next.isPanic() {
				return next
			}
			r = append(r, next)
		}
		return canonicalV(NewAV(r))
	default:
		return x
	}
}

func scan2Split(sep S, x V) V {
	switch xv := x.Value.(type) {
	case S:
		return NewAS(strings.Split(string(xv), string(sep)))
	case *AS:
		r := make([]V, xv.Len())
		for i := range r {
			r[i] = NewAS(strings.Split(xv.At(i), string(sep)))
		}
		return NewAV(r)
	case *AV:
		//assertCanonical(x)
		return panicf("s/x : x not a string atom or array (%s)", xv.Type())
	default:
		return panicf("s/x : x not a string atom or array (%s)", x.Type())
	}
}

func encodeBaseDigits(b int64, x int64) int {
	if b < 0 {
		b = -b
	}
	if x < 0 {
		x = -x
	}
	n := 1
	for x >= b {
		x /= b
		n++
	}
	return n
}

func scan2Encode(f V, x V) V {
	if f.IsI() {
		if f.I() == 0 {
			return panics("i\\x : base i is zero")
		}
		if x.IsI() {
			n := encodeBaseDigits(f.I(), x.I())
			r := make([]int64, n)
			for i := n - 1; i >= 0; i-- {
				r[i] = x.I() % f.I()
				x.N /= f.I()
			}
			return NewAI(r)
		}
		if x.IsF() {
			if !isI(x.F()) {
				return panicf("i\\x : x non-integer (%g)", x.F())
			}
			return scan2Encode(f, NewI(int64(x.F())))
		}
		switch xv := x.Value.(type) {
		case *AI:
			min, max := minMax(xv)
			max = maxI(absI(min), absI(max))
			n := encodeBaseDigits(f.I(), max)
			ai := make([]int64, n*xv.Len())
			copy(ai[(n-1)*xv.Len():], xv.Slice)
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < xv.Len(); j++ {
					ox := ai[i*xv.Len()+j]
					ai[i*xv.Len()+j] = ox % f.I()
					if i > 0 {
						ai[(i-1)*xv.Len()+j] = ox / f.I()
					}
				}
			}
			r := make([]V, n)
			for i := range r {
				r[i] = NewAI(ai[i*xv.Len() : (i+1)*xv.Len()])
			}
			return NewAV(r)
		case *AB:
			return scan2Encode(f, fromABtoAI(xv))
		case *AF:
			aix := toAI(xv)
			if aix.isPanic() {
				return aix
			}
			return scan2Encode(f, aix)
		case *AV:
			r := make([]V, xv.Len())
			for i, xi := range xv.Slice {
				r[i] = scan2Encode(f, xi)
				if r[i].isPanic() {
					return r[i]
				}
			}
			return canonicalV(NewAV(r))
		default:
			return panicType("i\\x", "x", x)
		}

	}
	if f.IsF() {
		if !isI(f.F()) {
			return panicf("i\\x : i non-integer (%g)", f.F())
		}
		return scan2Encode(NewI(int64(f.F())), x)
	}
	switch fv := f.Value.(type) {
	case *AB:
		return scan2Encode(fromABtoAI(fv), x)
	case *AI:
		if x.IsI() {
			// TODO: check for zero division
			n := fv.Len()
			r := make([]int64, n)
			for i := n - 1; i >= 0 && x.I() > 0; i-- {
				r[i] = x.I() % fv.At(i)
				x.N /= fv.At(i)
			}
			return NewAI(r)

		}
		if x.IsF() {
			if !isI(x.F()) {
				return panicf("I/x : x non-integer (%g)", x.F())
			}
			return scan2Encode(f, NewI(int64(x.F())))
		}
		switch xv := x.Value.(type) {
		case *AI:
			n := fv.Len()
			ai := make([]int64, n*xv.Len())
			copy(ai[(n-1)*xv.Len():], xv.Slice)
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < xv.Len(); j++ {
					ox := ai[i*xv.Len()+j]
					ai[i*xv.Len()+j] = ox % fv.At(i)
					if i > 0 {
						ai[(i-1)*xv.Len()+j] = ox / fv.At(i)
					}
				}
			}
			r := make([]V, n)
			for i := range r {
				r[i] = NewAI(ai[i*xv.Len() : (i+1)*xv.Len()])
			}
			return NewAV(r)
		case *AB:
			return scan2Encode(f, fromABtoAI(xv))
		case *AF:
			aix := toAI(xv)
			if aix.isPanic() {
				return aix
			}
			return scan2Encode(f, aix)
		case *AV:
			r := make([]V, xv.Len())
			for i, xi := range xv.Slice {
				r[i] = scan2Encode(f, xi)
				if r[i].isPanic() {
					return r[i]
				}
			}
			return canonicalV(NewAV(r))
		default:
			return panicType("I\\x", "x", x)
		}
	case *AF:
		aif := toAI(fv)
		if aif.isPanic() {
			return aif
		}
		return scan2Encode(aif, x)
	default:
		// should not happen
		return panicType("I\\x", "I", f)
	}
}

func scan3(ctx *Context, args []V) V {
	f := args[1]
	if !f.IsFunction() {
		return panicf("x f'y : f not a function (%s)", f.Type())
	}
	if f.Rank(ctx) != 2 {
		return scan3While(ctx, args)
	}
	y := args[0]
	x := args[2]
	switch yv := y.Value.(type) {
	case array:
		if yv.Len() == 0 {
			return NewAV([]V{})
		}
		ctx.push(yv.at(0))
		ctx.push(x)
		x.rcincr()
		first := ctx.applyN(f, 2)
		x.rcdecr()
		if first.isPanic() {
			return first
		}
		r := []V{first}
		for i := 1; i < yv.Len(); i++ {
			ctx.push(yv.at(i))
			last := r[len(r)-1]
			ctx.push(last)
			last.rcincr()
			next := ctx.applyN(f, 2)
			last.rcdecr()
			if next.isPanic() {
				return next
			}
			r = append(r, next)
		}
		return canonicalV(NewAV(r))
	default:
		ctx.push(y)
		ctx.push(x)
		return ctx.applyN(f, 2)
	}
}

func scan3While(ctx *Context, args []V) V {
	f := args[1]
	x := args[2]
	y := args[0]
	if x.IsI() {
		return scan3doTimes(ctx, x.I(), f, y)
	}
	if x.IsF() {
		if !isI(x.F()) {
			return panicf("n f\\y : non-integer n (%g)", x.F())
		}
		return scan3doTimes(ctx, int64(x.F()), f, y)
	}
	if x.IsFunction() {
		r := []V{y}
		for {
			ctx.push(y)
			y.rcincr()
			cond := ctx.applyN(x, 1)
			y.rcdecr()
			if cond.isPanic() {
				return cond
			}
			if !isTrue(cond) {
				return canonicalV(NewAV(r))
			}
			ctx.push(y)
			y = ctx.applyN(f, 1)
			if y.isPanic() {
				return y
			}
			r = append(r, y)
		}
	}
	return panicType("x f\\y", "x", x)
}

func scan3doTimes(ctx *Context, n int64, f, y V) V {
	r := []V{y}
	for i := int64(0); i < n; i++ {
		ctx.push(y)
		y = ctx.applyN(f, 1)
		if y.isPanic() {
			return y
		}
		r = append(r, y)
	}
	return canonicalV(NewAV(r))
}

func each2(ctx *Context, args []V) V {
	f := args[1]
	if !f.IsFunction() {
		return panicf("f'x : f not a function (%s)", f.Type())
	}
	x := toArray(args[0])
	switch xv := x.Value.(type) {
	case array:
		r := make([]V, 0, xv.Len())
		for i := 0; i < xv.Len(); i++ {
			ctx.push(xv.at(i))
			next := ctx.applyN(f, 1)
			if next.isPanic() {
				return next
			}
			r = append(r, next)
		}
		return canonicalV(NewAV(r))
	default:
		// should not happen
		return panicf("f'x : x not an array (%s)", x.Type())
	}
}

func each3(ctx *Context, args []V) V {
	f := args[1]
	if !f.IsFunction() {
		return panicf("x f'y : f not a function (%s)", f.Type())
	}
	x, okax := args[2].Value.(array)
	y, okay := args[0].Value.(array)
	if !okax && !okay {
		return ctx.ApplyN(f, args)
	}
	if !okax {
		ylen := y.Len()
		r := make([]V, 0, ylen)
		for i := 0; i < ylen; i++ {
			ctx.push(y.at(i))
			ctx.push(args[2])
			next := ctx.applyN(f, 2)
			if next.isPanic() {
				return next
			}
			r = append(r, next)
		}
		return canonicalV(NewAV(r))
	}
	if !okay {
		xlen := x.Len()
		r := make([]V, 0, xlen)
		for i := 0; i < xlen; i++ {
			ctx.push(args[0])
			ctx.push(x.at(i))
			next := ctx.applyN(f, 2)
			if next.isPanic() {
				return next
			}
			r = append(r, next)
		}
		return canonicalV(NewAV(r))
	}
	xlen := x.Len()
	if xlen != y.Len() {
		return panicf("x f'y : length mismatch: %d (#x) vs %d (#y)", x.Len(), y.Len())
	}
	r := make([]V, 0, xlen)
	for i := 0; i < xlen; i++ {
		ctx.push(y.at(i))
		ctx.push(x.at(i))
		next := ctx.applyN(f, 2)
		if next.isPanic() {
			return next
		}
		r = append(r, next)
	}
	return canonicalV(NewAV(r))
}
