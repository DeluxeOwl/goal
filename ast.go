package goal

import (
	"fmt"
	"strings"
)

// expr are built by the first left to right pass, resulting in a tree
// of blocks producing a whole expression, with simplified token
// information, and stack-like order. It is a non-resolved AST intermediary
// representation.
type expr interface {
	node()
}

// exprs is used to represent a whole expression. It is not a ppExpr.
type exprs []expr

// astToken represents a simplified token after processing into expr.
type astToken struct {
	Type astTokenType
	Rune rune
	Pos  int
	Text string
}

func (t *astToken) String() string {
	return fmt.Sprintf("{%v %d %s}", t.Type, t.Pos, t.Text)
}

// astTokenType represents tokens in a ppExpr.
type astTokenType int32

// These constants represent the possible tokens in a ppExpr. The SEP,
// EOF and CLOSE types are not emitted in the final result.
const (
	astSEP astTokenType = iota
	astEOF
	astCLOSE

	astNUMBER
	astSTRING
	astIDENT
	astVERB
	astADVERB
)

type astSpecial int

type astStrand []*astToken // for stranding, like 1 23 456

type astAdverbs []*astToken // for an adverb sequence

type astParenExpr exprs // for parenthesized sub-expressions

type astBlock struct {
	Type astBlockType
	Body []exprs
	Args []string
}

func (b *astBlock) String() (s string) {
	switch b.Type {
	case astLAMBDA:
		args := "[" + strings.Join([]string(b.Args), ";") + "]"
		s = fmt.Sprintf("{%s %v %v}", args, b.Type, b.Body)
	case astARGS:
		s = fmt.Sprintf("[%v %v]", b.Type, b.Body)
	case astLIST:
		s = fmt.Sprintf("(%v %v)", b.Type, b.Body)
	}
	return s
}

func (b *astBlock) push(e expr) {
	b.Body[len(b.Body)-1] = append(b.Body[len(b.Body)-1], e)
}

type astBlockType int

const (
	astLAMBDA astBlockType = iota
	astARGS
	astSEQ
	astLIST
)

func (t *astToken) node()    {}
func (st astStrand) node()   {}
func (ad astAdverbs) node()  {}
func (p astParenExpr) node() {}
func (b *astBlock) node()    {}

// astIter is an iterator for exprs slices, with peek functionality.
type astIter struct {
	exprs exprs
	i     int
}

func newAstIter(es exprs) astIter {
	return astIter{exprs: es, i: -1}
}

func (it *astIter) Next() bool {
	it.i++
	return it.i < len(it.exprs)
}

func (it *astIter) Expr() expr {
	return it.exprs[it.i]
}

func (it *astIter) Index() int {
	return it.i
}

func (it *astIter) Set(i int) {
	it.i = i
}

func (it *astIter) Peek() expr {
	if it.i+1 >= len(it.exprs) {
		return nil
	}
	return it.exprs[it.i+1]
}

func (it *astIter) PeekN(n int) expr {
	if it.i+n >= len(it.exprs) {
		return nil
	}
	return it.exprs[it.i+n]
}
