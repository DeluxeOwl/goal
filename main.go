package main

import (
	"fmt"
)

func main() {
	testVM("a:23 13;b:a+5;|b")
	testVM("a:!10;b:a+5;|b")
	testVM(`a:%0 2 0 3 4 5 2 2 2;a`)
	testVM(`a:=%"patata" "lolo" "patata" "patato";`)
	testVM(`a:=%"patata" "lolo" "patata" "patato"`)
	testVM(``)
	testVM(`a:1;b:{x+1} a`)
	testVM(`a:1;b:{x+y+2}[a;4]`)
	testVM(`a:1 3 5;b:3;a+b`)
	testVM(`a:1 3 5;f:{x+3};f[a]`)
	testVM(`a:1 3 5;;;|a`)
	testVM(`a:1 3 5;a[2 0 1 0]`)
	testVM(`(1;2;(3;4);4+1)`)
	testVM(`(1;2;(3;4);4+1;)`)
	testVM(`f:1+`)
	testVM(`f:1+;f 5`)
	testVM(`f:-+;f[5;2]`)
	testVM(`#(1;2;3)`)
	testVM(`#((2;3);(1;2;5))`)
	testVM(`#'((2;3);(1;2;5))`)
	testVM(`2 3#'1 2`)
	testVM(`{0 1 0 1} 0`)
	testVM(`{0 1 0 1}#1 2 3 4`)
	testVM(`+/!10`)
	testVM(`#0#1`)
	testVM(`+/0#1`)
	testVM(`+\!10`)
	//testVM(`+/!10000000`)
	testVM(`","/"a" "b" "c" "d"`)
	testVM(`-3 2`)
	testVM(`- 3 2`)
}

func testPrimitives() {
	fmt.Printf("Add:%#v\n", Add(I(3), I(5)))
	fmt.Printf("Add:%#v\n", Add(F(3), I(5)))
	fmt.Printf("Add:%#v\n", Add(I(3), AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AF{1, 2, 3}, AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AF{1, 2, 3, 4}, AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AV{AF{1, 2}, AF{3, 4}}, AI{3, 8}))
	fmt.Printf("Equal:%#v\n", Equal(AI{1, 3, 8, 2}, AI{5, 3, 8, 1}))
	//fmt.Printf("Add2:%#v\n", float64(1.0)/float64(0.0))
	fmt.Printf("Divide:%#v\n", Divide(F(2), F(0)))
	fmt.Printf("Sort:%#v\n", SortUp(AI{3, 2, 1}))
	fmt.Printf("Sort:%#v\n", SortUp(AV{I(3), I(2), I(1)}))
	fmt.Printf("Sort:%#v\n", SortUp(AV{I(3), I(2), AI{}, I(1), AI{2, 2}}))
	fmt.Printf("Take:%#v\n", Take(I(5), AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(I(2), AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(I(-5), AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(I(-2), AI{2, 3, 4}))
	fmt.Printf("ShiftBefore:%#v\n", ShiftBefore(AI{2, 3}, AI{1, 4, 5}))
	fmt.Printf("ShiftBefore:%#v\n", ShiftBefore(I(7), AF{1, 4, 5}))
	fmt.Printf("ShiftAfter:%#v\n", ShiftAfter(I(7), AF{1, 4, 5}))
	fmt.Printf("Flip:%#v\n", Flip(AI{1, 2, 3}))
	fmt.Printf("Flip:%#v\n", Flip(AV{AF{1}, AF{4}, AF{5}}))
	fmt.Printf("Flip:%#v\n", Flip(AV{AF{1, 2}, I(4), I(5)}))
	fmt.Printf("Flip:%#v\n", Flip(AV{AF{1, 2}, I(4), I(5), S("patata")}))
	fmt.Printf("Classify:%#v\n", Classify(AV{AF{1, 2}, I(4), AF{1, 2}, S("patata")}))
	fmt.Printf("Classify:%#v\n", Classify(AI{1, 2, 3, 2, 2, 4, 5, 3}))
	fmt.Printf("Range:%#v\n", Range(I(10)))
	fmt.Printf("Range:%#v\n", Range(AI{4, 2, 3}))
	fmt.Printf("Indices:%#v\n", Where(AI{0, 1, 0, 0, 1}))
	fmt.Printf("Indices:%#v\n", Where(AI{3, 0, 1}))
	fmt.Printf("MarkFirsts:%#v\n", MarkFirsts(AI{3, 3, 1, 2, 4, 2, 4}))
	fmt.Printf("OccurrenceCount:%#v\n", OccurrenceCount(AB{false, false, true, false, true, true}))
	fmt.Printf("Windows:%#v\n", Windows(I(3), Range(I(7))))
	fmt.Printf("MemberOf:%#v\n", MemberOf(AS{"two", "twelve", "five", "one", "one", "nine"}, AS{"one", "two", "four"}))
	fmt.Printf("MemberOf:%#v\n", MemberOf(I(5), AI{2, 3, 6}))
	fmt.Printf("MemberOf:%#v\n", MemberOf(I(3), AV{I(2), I(3), I(6)}))
	fmt.Printf("MemberOf:%#v\n", MemberOf(AI{4, 3, 3, 3, 5, 2, 6}, AI{2, 3, 6}))
	fmt.Printf("Group:%#v\n", Group(AI{0, 3, 2, 2, 0, 3}))
	fmt.Printf("Group:%#v\n", Group(AB{false, true, false}))
}

func testVM(s string) {
	fmt.Println("-------- Goal code ----------")
	fmt.Println(s)
	ctx := NewContext()
	v, err := ctx.RunString(s)
	ctx.Show()
	if err != nil {
		fmt.Println("---------- Error -----------")
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("---------- Result -----------")
	fmt.Printf("%v\n", v)
}
