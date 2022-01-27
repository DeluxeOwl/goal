package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, world!\n")
	fmt.Printf("Add:%#v\n", Add(I(3), I(5)))
	fmt.Printf("Add:%#v\n", Add(F(3), I(5)))
	fmt.Printf("Add:%#v\n", Add(I(3), AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AF{1, 2, 3}, AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AF{1, 2, 3, 4}, AI{5, 3, 8}))
	fmt.Printf("Add:%#v\n", Add(AO{AF{1, 2}, AF{3, 4}}, AI{3, 8}))
	fmt.Printf("Equal:%#v\n", Equal(AI{1, 3, 8, 2}, AI{5, 3, 8, 1}))
	//fmt.Printf("Add2:%#v\n", float64(1.0)/float64(0.0))
	fmt.Printf("Divide:%#v\n", Divide(F(2), F(0)))
	fmt.Printf("Sort:%#v\n", SortUp(AI{3, 2, 1}))
	fmt.Printf("Sort:%#v\n", SortUp(AO{3, 2, 1}))
	fmt.Printf("Sort:%#v\n", SortUp(AO{3, 2, AI{}, 1, AI{2, 2}}))
	fmt.Printf("Take:%#v\n", Take(5, AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(2, AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(-5, AI{2, 3, 4}))
	fmt.Printf("Take:%#v\n", Take(-2, AI{2, 3, 4}))
}
