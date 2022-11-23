package goal

import "strings"

// Matcher is implemented by types that can be matched againts other objects
// (typically a struct of the same type with fields that match).
type Matcher interface {
	Matches(x V) bool
}

// Match returns true if the two values match like in x~y.
func Match(x, y V) bool {
	return x != nil && x.Matches(y) || x == nil && y == nil
}

func matchArray(x array, y V) bool {
	ya, ok := y.(array)
	if !ok {
		return false
	}
	l := x.Len()
	if l != ya.Len() {
		return false
	}
	switch x := x.(type) {
	case AB:
		switch ya := ya.(type) {
		case AB:
			return matchAB(x, ya)
		case AI:
			return matchABAI(x, ya)
		case AF:
			return matchABAF(x, ya)
		}
	case AI:
		switch ya := ya.(type) {
		case AB:
			return matchABAI(ya, x)
		case AI:
			return matchAI(x, ya)
		case AF:
			return matchAIAF(x, ya)
		}
	case AF:
		switch ya := ya.(type) {
		case AB:
			return matchABAF(ya, x)
		case AI:
			return matchAIAF(ya, x)
		case AF:
			return matchAF(x, ya)
		}
	case AS:
		ya, ok := ya.(AS)
		if !ok {
			break
		}
		for i, yi := range ya {
			if yi != x[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < l; i++ {
		if !Match(x.at(i), ya.at(i)) {
			return false
		}
	}
	return true
}

func matchAB(x, y AB) bool {
	for i, yi := range y {
		if yi != x[i] {
			return false
		}
	}
	return true
}

func matchABAI(x AB, y AI) bool {
	for i, yi := range y {
		if yi != int(B2I(x[i])) {
			return false
		}
	}
	return true
}

func matchABAF(x AB, y AF) bool {
	for i, yi := range y {
		if F(yi) != B2F(x[i]) {
			return false
		}
	}
	return true
}

func matchAI(x, y AI) bool {
	for i, yi := range y {
		if yi != x[i] {
			return false
		}
	}
	return true
}

func matchAIAF(x AI, y AF) bool {
	for i, yi := range y {
		if F(yi) != F(x[i]) {
			return false
		}
	}
	return true
}

func matchAF(x, y AF) bool {
	for i, yi := range y {
		if yi != x[i] {
			return false
		}
	}
	return true
}

// classify returns %x.
func classify(x V) V {
	if Length(x) == 0 {
		return AB{}
	}
	assertCanonical(x)
	switch x := x.(type) {
	case F, I, S:
		return errf("%%x : x not an array (%s)", x.Type())
	case AB:
		if !x[0] {
			return x
		}
		return not(x)
	case AF:
		r := make(AI, len(x))
		m := map[float64]int{}
		n := 0
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				r[i] = n
				m[xi] = n
				n++
				continue
			}
			r[i] = c
		}
		return r
	case AI:
		r := make(AI, len(x))
		m := map[int]int{}
		n := 0
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				r[i] = n
				m[xi] = n
				n++
				continue
			}
			r[i] = c
		}
		return r
	case AS:
		r := make(AI, len(x))
		m := map[string]int{}
		n := 0
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				r[i] = n
				m[xi] = n
				n++
				continue
			}
			r[i] = c
		}
		return r
	case AV:
		// NOTE: quadratic algorithm, worst case complexity could be
		// improved by sorting or string hashing.
		r := make(AI, len(x))
		n := 0
	loop:
		for i, xi := range x {
			for j := range x[:i] {
				if Match(xi, x[j]) {
					r[i] = r[j]
					continue loop
				}
			}
			r[i] = n
			n++
		}
		return r
	default:
		return errf("%%x : x not an array (%s)", x.Type())
	}
}

// uniq returns ?x.
func uniq(x V) V {
	if Length(x) == 0 {
		return x
	}
	assertCanonical(x)
	switch x := x.(type) {
	case F, I, S:
		// NOTE: ?atom could be used for something.
		return errf("?x : x not an array (%s)", x.Type())
	case AB:
		if len(x) == 0 {
			return x
		}
		b := x[0]
		for i := 1; i < len(x); i++ {
			if x[i] != b {
				return AB{b, x[i]}
			}
		}
		return AB{b}
	case AF:
		r := AF{}
		m := map[float64]struct{}{}
		for _, xi := range x {
			_, ok := m[xi]
			if !ok {
				r = append(r, xi)
				m[xi] = struct{}{}
			}
		}
		return r
	case AI:
		r := AI{}
		m := map[int]struct{}{}
		for _, xi := range x {
			_, ok := m[xi]
			if !ok {
				r = append(r, xi)
				m[xi] = struct{}{}
				continue
			}
		}
		return r
	case AS:
		r := AS{}
		m := map[string]struct{}{}
		for _, xi := range x {
			_, ok := m[xi]
			if !ok {
				r = append(r, xi)
				m[xi] = struct{}{}
				continue
			}
		}
		return r
	case AV:
		// NOTE: quadratic algorithm, worst case complexity could be
		// improved by sorting or string hashing.
		r := make(AV, len(x))
	loop:
		for i, xi := range x {
			for j := range x[:i] {
				if Match(xi, x[j]) {
					continue loop
				}
			}
			r = append(r, xi)
		}
		return canonical(r)
	default:
		return errf("?x : x not an array (%s)", x.Type())
	}
}

// Mark Firsts returns ∊x. XXX unused for now
func markFirsts(x V) V {
	if Length(x) == 0 {
		return AB{}
	}
	assertCanonical(x)
	switch x := x.(type) {
	case F, I, S:
		return errf("∊x : x not an array (%s)", x.Type())
	case AB:
		r := make(AB, len(x))
		r[0] = true
		x0 := x[0]
		for i := 1; i < len(x); i++ {
			if x[i] != x0 {
				r[i] = true
				break
			}
		}
		return r
	case AF:
		r := make(AB, len(x))
		m := map[float64]struct{}{}
		for i, xi := range x {
			_, ok := m[xi]
			if !ok {
				r[i] = true
				m[xi] = struct{}{}
				continue
			}
		}
		return r
	case AI:
		r := make(AB, len(x))
		m := map[int]struct{}{}
		for i, xi := range x {
			_, ok := m[xi]
			if !ok {
				r[i] = true
				m[xi] = struct{}{}
				continue
			}
		}
		return r
	case AS:
		r := make(AB, len(x))
		m := map[string]struct{}{}
		for i, xi := range x {
			_, ok := m[xi]
			if !ok {
				r[i] = true
				m[xi] = struct{}{}
				continue
			}
		}
		return r
	case AV:
		// NOTE: quadratic algorithm, worst case complexity could be
		// improved by sorting or string hashing.
		r := make(AB, len(x))
	loop:
		for i, xi := range x {
			for j := range x[:i] {
				if Match(xi, x[j]) {
					continue loop
				}
			}
			r[i] = true
		}
		return r
	default:
		return errf("∊x : x not an array (%s)", x.Type())
	}
}

// memberOf returns x in y.
func memberOf(x, y V) V {
	if Length(y) == 0 {
		switch x := x.(type) {
		case array:
			r := make(AB, x.Len())
			return r
		default:
			return B2I(false)
		}
	}
	if Length(x) == 0 {
		return AB{}
	}
	assertCanonical(x)
	assertCanonical(y)
	switch y := y.(type) {
	case AB:
		return memberOfAB(x, y)
	case AF:
		return memberOfAF(x, y)
	case AI:
		return memberOfAI(x, y)
	case AS:
		return memberOfAS(x, y)
	case AV:
		return memberOfAV(x, y)
	default:
		return errf("x in y : y not an array (%s)", y.Type())
	}
}

func memberOfAB(x V, y AB) V {
	var t, f bool
	for _, yi := range y {
		if t && f {
			break
		}
		t, f = t || yi, f || !yi
	}
	if t && f {
		switch x := x.(type) {
		case array:
			r := make(AB, x.Len())
			for i := range r {
				r[i] = true
			}
			return r
		default:
			return B2I(true)
		}
	}
	if t {
		return equal(x, B2I(true))
	}
	return equal(x, B2I(false))
}

func memberOfAF(x V, y AF) V {
	m := map[F]struct{}{}
	for _, yi := range y {
		_, ok := m[F(yi)]
		if !ok {
			m[F(yi)] = struct{}{}
			continue
		}
	}
	switch x := x.(type) {
	case I:
		_, ok := m[F(x)]
		return B2I(ok)
	case F:
		_, ok := m[x]
		return B2I(ok)
	case AB:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[B2F(xi)]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[F(xi)]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[F(xi)]
		}
		return r
	case array:
		return memberOfArray(x, y)
	default:
		return make(AB, Length(x))
	}
}

func memberOfAI(x V, y AI) V {
	m := map[int]struct{}{}
	for _, yi := range y {
		_, ok := m[yi]
		if !ok {
			m[yi] = struct{}{}
			continue
		}
	}
	switch x := x.(type) {
	case I:
		_, ok := m[int(x)]
		return B2I(ok)
	case F:
		if !isI(x) {
			return B2I(false)
		}
		_, ok := m[int(x)]
		return B2I(ok)
	case AB:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[int(B2I(xi))]
		}
		return r
	case AI:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[xi]
		}
		return r
	case AF:
		r := make(AB, len(x))
		for i, xi := range x {
			if !isI(F(xi)) {
				continue
			}
			_, r[i] = m[int(xi)]
		}
		return r
	case array:
		return memberOfArray(x, y)
	default:
		return make(AB, Length(x))
	}
}

func memberOfAS(x V, y AS) V {
	m := map[string]struct{}{}
	for _, yi := range y {
		_, ok := m[yi]
		if !ok {
			m[yi] = struct{}{}
			continue
		}
	}
	switch x := x.(type) {
	case S:
		_, ok := m[string(x)]
		return B2I(ok)
	case AS:
		r := make(AB, len(x))
		for i, xi := range x {
			_, r[i] = m[xi]
		}
		return r
	case array:
		return memberOfArray(x, y)
	default:
		return make(AB, Length(x))
	}
}

func memberOfAV(x V, y AV) V {
	switch x := x.(type) {
	case array:
		return memberOfArray(x, y)
	default:
		for _, yi := range y {
			if Match(x, yi) {
				return B2I(true)
			}
		}
		return B2I(false)
	}
}

func memberOfArray(x, y array) V {
	// NOTE: quadratic algorithm, worst case complexity could be
	// improved by sorting or string hashing.
	res := make(AB, x.Len())
	for i := 0; i < x.Len(); i++ {
		for j := 0; j < y.Len(); j++ {
			if Match(x.at(i), y.at(j)) {
				res[i] = true
				break
			}
		}
	}
	return res
}

// OccurrenceCount returns ⊒x.
func occurrenceCount(x V) V {
	if Length(x) == 0 {
		return AB{}
	}
	assertCanonical(x)
	switch x := x.(type) {
	case F, I, S:
		return errf("⊒x : x not an array (%s)", x.Type())
	case AB:
		r := make(AI, len(x))
		var f, t int
		for i, xi := range x {
			if xi {
				r[i] = t
				t++
				continue
			}
			r[i] = f
			f++
		}
		return r
	case AF:
		r := make(AI, len(x))
		m := map[float64]int{}
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				m[xi] = 0
				continue
			}
			m[xi] = c + 1
			r[i] = c + 1
		}
		return r
	case AI:
		r := make(AI, len(x))
		m := map[int]int{}
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				m[xi] = 0
				continue
			}
			m[xi] = c + 1
			r[i] = c + 1
		}
		return r
	case AS:
		r := make(AI, len(x))
		m := map[string]int{}
		for i, xi := range x {
			c, ok := m[xi]
			if !ok {
				m[xi] = 0
				continue
			}
			m[xi] = c + 1
			r[i] = c + 1
		}
		return r
	case AV:
		// NOTE: quadratic algorithm, worst case complexity could be
		// improved by sorting or string hashing.
		r := make(AI, len(x))
	loop:
		for i, xi := range x {
			for j := i - 1; j >= 0; j-- {
				if Match(xi, x[j]) {
					r[i] = r[j] + 1
					continue loop
				}
			}
		}
		return r
	default:
		return errf("⊒x : x not an array (%s)", x.Type())
	}
}

// without returns x^y.
func without(x, y V) V {
	switch z := x.(type) {
	case I:
		return windows(int(z), y)
	case F:
		if !isI(z) {
			return errf("i^y : i non-integer (%g)", z)
		}
		return windows(int(z), y)
	case S:
		return trim(z, y)
	case array:
		y = toArray(y)
		res := memberOf(y, x)
		switch bres := res.(type) {
		case I:
			res = I(1 - bres)
		case AB:
			for i, b := range bres {
				bres[i] = !b
			}
		}
		res = replicate(res, y)
		return res
	default:
		return errType("x^y", "x", x)
	}
}

// find returns x?y.
func find(x, y V) V {
	assertCanonical(y)
	assertCanonical(x)
	switch x := x.(type) {
	case S:
		return findS(x, y)
	case AB:
		return findAB(x, y)
	case AF:
		return findAF(x, y)
	case AI:
		return findAI(x, y)
	case AS:
		return findAS(x, y)
	case AV:
		return findAV(x, y)
	default:
		return errf("x?y : x not an array (%s)", x.Type())
	}
}

func findS(s S, y V) V {
	switch y := y.(type) {
	case S:
		return I(strings.Index(string(s), string(y)))
	case AS:
		res := make(AI, y.Len())
		for i, ss := range y {
			res[i] = strings.Index(string(s), string(ss))
		}
		return res
	case AV:
		res := make(AV, y.Len())
		for i, yi := range y {
			res[i] = findS(s, yi)
			if err, ok := res[i].(errV); ok {
				return err
			}
		}
		return res
	default:
		return errType("s?y", "y", y)
	}
}

func imapAB(x AB) (m [2]int) {
	m[0] = len(x)
	m[1] = len(x)
	if len(x) == 0 {
		return m
	}
	m[int(B2I(x[0]))] = 0
	for i, xi := range x[1:] {
		if xi != x[0] {
			m[int(B2I(xi))] = i + 1
			break
		}
	}
	return m
}

func imapAI(x AI) map[int]int {
	m := map[int]int{}
	for i, xi := range x {
		_, ok := m[xi]
		if !ok {
			m[xi] = i
			continue
		}
	}
	return m
}

func imapAF(x AF) map[float64]int {
	m := map[float64]int{}
	for i, xi := range x {
		_, ok := m[xi]
		if !ok {
			m[xi] = i
			continue
		}
	}
	return m
}

func imapAS(x AS) map[string]int {
	m := map[string]int{}
	for i, xi := range x {
		_, ok := m[xi]
		if !ok {
			m[xi] = i
			continue
		}
	}
	return m
}

func findAB(x AB, y V) V {
	switch y := y.(type) {
	case I:
		for i, xi := range x {
			if B2I(xi) == y {
				return I(i)
			}
		}
		return I(x.Len())
	case F:
		if !isI(y) {
			return I(x.Len())
		}
		return findAB(x, I(y))
	case AB:
		m := imapAB(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			res[i] = m[B2I(yi)]
		}
		return res
	case AI:
		m := imapAB(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			if yi != 0 && yi != 1 {
				res[i] = x.Len()
			} else {
				res[i] = m[yi]
			}
		}
		return res
	case AF:
		m := imapAB(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			if yi != 0 && yi != 1 {
				res[i] = x.Len()
			} else {
				res[i] = m[int(yi)]
			}
		}
		return res
	case array:
		return findArray(x, y)
	default:
		return I(x.Len())
	}
}

func findAF(x AF, y V) V {
	switch y := y.(type) {
	case I:
		for i, xi := range x {
			if xi == float64(y) {
				return I(i)
			}
		}
		return I(x.Len())
	case F:
		for i, xi := range x {
			if F(xi) == y {
				return I(i)
			}
		}
		return I(x.Len())
	case AB:
		m := imapAF(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[float64(B2F(yi))]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case AI:
		m := imapAF(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[float64(yi)]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case AF:
		m := imapAF(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[yi]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case array:
		return findArray(x, y)
	default:
		return I(x.Len())
	}
}

func findAI(x AI, y V) V {
	switch y := y.(type) {
	case I:
		for i, xi := range x {
			if I(xi) == y {
				return I(i)
			}
		}
		return I(x.Len())
	case F:
		for i, xi := range x {
			if F(xi) == y {
				return I(i)
			}
		}
		return I(x.Len())
	case AB:
		m := imapAI(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[int(B2I(yi))]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case AI:
		m := imapAI(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[yi]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case AF:
		m := imapAI(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			if !isI(F(yi)) {
				res[i] = x.Len()
				continue
			}
			j, ok := m[int(yi)]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case array:
		return findArray(x, y)
	default:
		return I(x.Len())
	}
}

func findAS(x AS, y V) V {
	switch y := y.(type) {
	case S:
		for i, xi := range x {
			if S(xi) == y {
				return I(i)
			}
		}
		return I(x.Len())
	case AS:
		m := imapAS(x)
		res := make(AI, y.Len())
		for i, yi := range y {
			j, ok := m[yi]
			if ok {
				res[i] = j
			} else {
				res[i] = x.Len()
			}
		}
		return res
	case array:
		return findArray(x, y)
	default:
		return I(x.Len())
	}
}

func findArray(x, y array) V {
	// NOTE: quadratic algorithm, worst case complexity could be
	// improved by sorting or string hashing.
	res := make(AI, y.Len())
	for i := range res {
		res[i] = x.Len()
	}
	for i := 0; i < y.Len(); i++ {
		for j := 0; j < x.Len(); j++ {
			if Match(y.at(i), x.at(j)) {
				res[i] = j
				break
			}
		}
	}
	return res
}

func findAV(x AV, y V) V {
	switch y := y.(type) {
	case F:
		for i, xi := range x {
			if Match(xi, y) {
				return I(i)
			}
		}
		return I(x.Len())
	case I:
		for i, xi := range x {
			if Match(xi, y) {
				return I(i)
			}
		}
		return I(x.Len())
	case S:
		for i, xi := range x {
			if Match(xi, y) {
				return I(i)
			}
		}
		return I(x.Len())
	case array:
		return findArray(x, y)
	default:
		for i, xi := range x {
			if Match(xi, y) {
				return I(i)
			}
		}
		return I(x.Len())
	}
}
