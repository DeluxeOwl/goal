package goal

import (
	"sort"
	"strings"
)

type sortAB []bool

func (bs sortAB) Len() int {
	return len(bs)
}

func (bs sortAB) Less(i, j int) bool {
	return bs[j] && !bs[i]
}

func (bs sortAB) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

type sortVSlice []V

func (bs sortVSlice) Len() int {
	return len(bs)
}

func (bs sortVSlice) Less(i, j int) bool {
	return less(bs[i], bs[j])
}

func (bs sortVSlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func less(x, y V) bool {
	if x.IsInt() {
		return lessI(x, y)
	}
	switch xv := x.Value.(type) {
	case F:
		return lessF(x, y)
	case S:
		return lessS(x, y)
	case *AB:
		if xv.Len() == 0 {
			return Length(y) > 0
		}
		return lessAB(x, y)
	case *AF:
		if xv.Len() == 0 {
			return Length(y) > 0
		}
		return lessAF(x, y)
	case *AI:
		if xv.Len() == 0 {
			return Length(y) > 0
		}
		return lessAI(x, y)
	case *AS:
		if xv.Len() == 0 {
			return Length(y) > 0
		}
		return lessAS(x, y)
	case *AV:
		if xv.Len() == 0 {
			return Length(y) > 0
		}
		return lessAV(x, y)
	default:
		return false
	}
}

func lessF(x V, y V) bool {
	xv := x.Value.(F)
	if y.IsInt() {
		return xv < F(y.Int())
	}
	switch yv := y.Value.(type) {
	case F:
		return xv < yv
	case *AB:
		if yv.Len() == 0 {
			return false
		}
		return xv < B2F(yv.At(0)) || xv == B2F(yv.At(0)) && yv.Len() > 1
	case *AF:
		if yv.Len() == 0 {
			return false
		}
		return xv < F(yv.At(0)) || xv == F(yv.At(0)) && yv.Len() > 1
	case *AI:
		if yv.Len() == 0 {
			return false
		}
		return xv < F(yv.At(0)) || xv == F(yv.At(0)) && yv.Len() > 1
	case *AV:
		if yv.Len() == 0 {
			return false
		}
		return lessF(x, yv.At(0)) || !less(yv.At(0), x) && yv.Len() > 1
	default:
		return false
	}
}

func lessI(x V, y V) bool {
	xv := x.Int()
	if y.IsInt() {
		return xv < y.Int()
	}
	switch yv := y.Value.(type) {
	case F:
		return F(xv) < yv
	case *AB:
		if yv.Len() == 0 {
			return false
		}
		return xv < B2I(yv.At(0)) || xv == B2I(yv.At(0)) && yv.Len() > 1
	case *AF:
		if yv.Len() == 0 {
			return false
		}
		return float64(xv) < yv.At(0) || float64(xv) == yv.At(0) && yv.Len() > 1
	case *AI:
		if yv.Len() == 0 {
			return false
		}
		return xv < yv.At(0) || xv == yv.At(0) && yv.Len() > 1
	case *AV:
		if yv.Len() == 0 {
			return false
		}
		return lessI(x, yv.At(0)) || !less(yv.At(0), x) && yv.Len() > 1
	default:
		return false
	}
}

func lessS(x V, y V) bool {
	xv := x.Value.(S)
	switch yv := y.Value.(type) {
	case S:
		return xv < yv
	case *AS:
		if yv.Len() == 0 {
			return false
		}
		return string(xv) < yv.At(0) || string(xv) == yv.At(0) && yv.Len() > 1
	case *AV:
		if yv.Len() == 0 {
			return false
		}
		return lessS(x, yv.At(0)) || !less(yv.At(0), x) && yv.Len() > 1
	default:
		return false
	}
}

func lessAB(x V, y V) bool {
	xv := x.Value.(*AB)
	if y.IsInt() {
		return !lessI(y, x)
	}
	switch yv := y.Value.(type) {
	case F:
		return !lessF(y, x)
	case *AB:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) && !yv.At(i) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AF:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if B2F(xv.At(i)) > F(yv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AI:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if B2I(xv.At(i)) > yv.At(i) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AV:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(yv.At(i), NewI(B2I(xv.At(i)))) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	default:
		return false
	}
}

func lessAI(x V, y V) bool {
	xv := x.Value.(*AI)
	if y.IsInt() {
		return !lessI(y, x)
	}
	switch yv := y.Value.(type) {
	case F:
		return !lessF(y, x)
	case *AB:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) > B2I(yv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AF:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if F(xv.At(i)) > F(yv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AI:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) > yv.At(i) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AV:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(yv.At(i), NewI(xv.At(i))) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	default:
		return false
	}
}

func lessAF(x V, y V) bool {
	xv := x.Value.(*AF)
	if y.IsInt() {
		return !lessI(y, x)
	}
	switch yv := y.Value.(type) {
	case F:
		return !lessF(y, x)
	case *AB:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if F(xv.At(i)) > B2F(yv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AF:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) > yv.At(i) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AI:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) > float64(yv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AV:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(yv.At(i), NewF(xv.At(i))) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	default:
		return false
	}
}

func lessAS(x V, y V) bool {
	xv := x.Value.(*AS)
	switch yv := y.Value.(type) {
	case S:
		return !lessS(y, x)
	case *AS:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if xv.At(i) > yv.At(i) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AV:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(yv.At(i), NewS(xv.At(i))) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	default:
		return false
	}
}

func lessAV(x V, y V) bool {
	xv := x.Value.(*AV)
	if y.IsInt() {
		return less(xv.At(0), y)
	}
	switch yv := y.Value.(type) {
	case F:
		return less(xv.At(0), y)
	case *AB:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(NewI(B2I(yv.At(i))), xv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AF:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(NewF(yv.At(i)), xv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AI:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(NewI(yv.At(i)), xv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	case *AV:
		for i := 0; i < xv.Len() && i < yv.Len(); i++ {
			if less(yv.At(i), xv.At(i)) {
				return false
			}
		}
		return xv.Len() < yv.Len()
	default:
		return false
	}
}

// sortUp returns ^x.
func sortUp(x V) V {
	// TODO: avoid cases of double clone
	//assertCanonical(x)
	x = cloneShallow(x)
	switch x := x.Value.(type) {
	case *AB:
		sort.Stable(sortAB(x.Slice))
		return NewV(x)
	case *AF:
		sort.Stable(sort.Float64Slice(x.Slice))
		return NewV(x)
	case *AI:
		sort.Stable(sort.IntSlice(x.Slice))
		return NewV(x)
	case *AS:
		sort.Stable(sort.StringSlice(x.Slice))
		return NewV(x)
	case *AV:
		sort.Stable(sortVSlice(x.Slice))
		return NewV(x)
	default:
		return errf("^x : x not an array (%s)", x.Type())
	}
}

type permutationAV struct {
	Perm []int
	X    sortVSlice
}

func (p *permutationAV) Len() int {
	return p.X.Len()
}

func (p *permutationAV) Swap(i, j int) {
	p.Perm[i], p.Perm[j] = p.Perm[j], p.Perm[i]
}

func (p *permutationAV) Less(i, j int) bool {
	return p.X.Less(p.Perm[i], p.Perm[j])
}

type permutationAB struct {
	Perm []int
	X    sortAB
}

func (p *permutationAB) Len() int {
	return p.X.Len()
}

func (p *permutationAB) Swap(i, j int) {
	p.Perm[i], p.Perm[j] = p.Perm[j], p.Perm[i]
}

func (p *permutationAB) Less(i, j int) bool {
	return p.X.Less(p.Perm[i], p.Perm[j])
}

type permutationAI struct {
	Perm []int
	X    sort.IntSlice
}

func (p *permutationAI) Len() int {
	return p.X.Len()
}

func (p *permutationAI) Swap(i, j int) {
	p.Perm[i], p.Perm[j] = p.Perm[j], p.Perm[i]
}

func (p *permutationAI) Less(i, j int) bool {
	return p.X.Less(p.Perm[i], p.Perm[j])
}

type permutationAF struct {
	Perm []int
	X    sort.Float64Slice
}

func (p *permutationAF) Len() int {
	return p.X.Len()
}

func (p *permutationAF) Swap(i, j int) {
	p.Perm[i], p.Perm[j] = p.Perm[j], p.Perm[i]
}

func (p *permutationAF) Less(i, j int) bool {
	return p.X.Less(p.Perm[i], p.Perm[j])
}

type permutationAS struct {
	Perm []int
	X    sort.StringSlice
}

func (p *permutationAS) Len() int {
	return p.X.Len()
}

func (p *permutationAS) Swap(i, j int) {
	p.Perm[i], p.Perm[j] = p.Perm[j], p.Perm[i]
}

func (p *permutationAS) Less(i, j int) bool {
	return p.X.Less(p.Perm[i], p.Perm[j])
}

func permRange(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

// ascend returns <x.
func ascend(x V) V {
	switch x := x.Value.(type) {
	case *AB:
		p := &permutationAB{Perm: permRange(x.Len()), X: sortAB(x.Slice)}
		sort.Stable(p)
		return NewAI(p.Perm)
	case *AF:
		p := &permutationAF{Perm: permRange(x.Len()), X: sort.Float64Slice(x.Slice)}
		sort.Stable(p)
		return NewAI(p.Perm)
	case *AI:
		p := &permutationAI{Perm: permRange(x.Len()), X: sort.IntSlice(x.Slice)}
		sort.Stable(p)
		return NewAI(p.Perm)
	case *AS:
		p := &permutationAS{Perm: permRange(x.Len()), X: sort.StringSlice(x.Slice)}
		sort.Stable(p)
		return NewAI(p.Perm)
	case *AV:
		p := &permutationAV{Perm: permRange(x.Len()), X: sortVSlice(x.Slice)}
		sort.Stable(p)
		return NewAI(p.Perm)
	default:
		return errf("<x : x not an array (%s)", x.Type())
	}
}

// descend returns >x.
func descend(x V) V {
	p := ascend(x)
	if p.IsErr() {
		return errs(">" + strings.TrimPrefix(p.Value.(errV).Error(), "<"))
	}
	reverseMut(p)
	return p
}

// search implements x$y.
func search(x V, y V) V {
	switch xv := x.Value.(type) {
	case *AB:
		if !sort.IsSorted(sortAB(xv.Slice)) {
			return errDomain("x$y", "x is not ascending")
		}
		return searchAI(fromABtoAI(xv).Value.(*AI), y)
	case *AI:
		if !sort.IsSorted(sort.IntSlice(xv.Slice)) {
			return errDomain("x$y", "x is not ascending")
		}
		return searchAI(xv, y)
	case *AF:
		if !sort.IsSorted(sort.Float64Slice(xv.Slice)) {
			return errDomain("x$y", "x is not ascending")
		}
		return searchAF(xv, y)
	case *AS:
		if !sort.IsSorted(sort.StringSlice(xv.Slice)) {
			return errDomain("x$y", "x is not ascending")
		}
		return searchAS(xv, y)
	case *AV:
		if !sort.IsSorted(sortVSlice(xv.Slice)) {
			return errDomain("x$y", "x is not ascending")
		}
		return searchAV(xv, y)
	default:
		// should not happen
		return errType("x$y", "x", x)
	}
}

func searchAII(x *AI, y int) int {
	return sort.Search(x.Len(), func(i int) bool { return x.At(i) > y })
}

func searchAIF(x *AI, y F) int {
	return sort.Search(x.Len(), func(i int) bool { return F(x.At(i)) > y })
}

func searchAFI(x *AF, y int) int {
	return sort.Search(x.Len(), func(i int) bool { return x.At(i) > float64(y) })
}

func searchAFF(x *AF, y F) int {
	return sort.Search(x.Len(), func(i int) bool { return F(x.At(i)) > y })
}

func searchASS(x *AS, y S) int {
	return sort.Search(x.Len(), func(i int) bool { return S(x.At(i)) > y })
}

func searchAI(x *AI, y V) V {
	if y.IsInt() {
		return NewI(searchAII(x, y.Int()))
	}
	switch y := y.Value.(type) {
	case F:
		return NewI(searchAIF(x, y))
	case *AB:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAII(x, B2I(yi))
		}
		return NewAI(r)
	case *AI:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAII(x, yi)
		}
		return NewAI(r)
	case *AF:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAIF(x, F(yi))
		}
		return NewAI(r)
	case array:
		r := make([]int, y.Len())
		for i := 0; i < y.Len(); i++ {
			r[i] = sort.Search(x.Len(),
				func(i int) bool { return less(y.at(i), NewI(x.At(i))) })
		}
		return NewAI(r)
	default:
		return NewI(x.Len())
	}
}

func searchAF(x *AF, y V) V {
	if y.IsInt() {
		return NewI(searchAFI(x, y.Int()))
	}
	switch y := y.Value.(type) {
	case F:
		return NewI(searchAFF(x, y))
	case *AB:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAFI(x, B2I(yi))
		}
		return NewAI(r)
	case *AI:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAFI(x, yi)
		}
		return NewAI(r)
	case *AF:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchAFF(x, F(yi))
		}
		return NewAI(r)
	case array:
		r := make([]int, y.Len())
		for i := 0; i < y.Len(); i++ {
			r[i] = sort.Search(x.Len(),
				func(i int) bool { return less(y.at(i), NewF(x.At(i))) })
		}
		return NewAI(r)
	default:
		return NewI(x.Len())
	}
}

func searchAS(x *AS, y V) V {
	switch y := y.Value.(type) {
	case S:
		return NewI(searchASS(x, y))
	case *AS:
		r := make([]int, y.Len())
		for i, yi := range y.Slice {
			r[i] = searchASS(x, S(yi))
		}
		return NewAI(r)
	case array:
		r := make([]int, y.Len())
		for i := 0; i < y.Len(); i++ {
			r[i] = sort.Search(x.Len(),
				func(i int) bool { return less(y.at(i), NewS(x.At(i))) })
		}
		return NewAI(r)
	default:
		return NewI(x.Len())
	}
}

func searchAV(x *AV, y V) V {
	switch yv := y.Value.(type) {
	case array:
		r := make([]int, yv.Len())
		for i := 0; i < yv.Len(); i++ {
			r[i] = sort.Search(x.Len(),
				func(i int) bool { return less(yv.at(i), x.At(i)) })
		}
		return NewAI(r)
	default:
		return NewI(sort.Search(x.Len(),
			func(i int) bool { return less(y, x.At(i)) }))

	}
}
