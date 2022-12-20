package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Value struct {
	number      int
	moved       bool
	left, right *Value
}

func (v *Value) Move(newLeft, newRight *Value) {
	// connect neighbors of subject's old position
	v.left.right = v.right
	v.right.left = v.left

	// update subjects neighbors
	v.left = newLeft
	v.right = newRight

	// update new neighbors
	newLeft.right = v
	newRight.left = v
}

func (v Value) RightNeighborN(n int) *Value {
	current := &v

	for i := 0; i < n; i++ {
		current = current.right
	}
	return current
}

func (v Value) LeftNeighborN(n int) *Value {
	current := &v

	for ; n <= 0; n++ {
		current = current.left
	}
	return current
}

func (v Value) GetNeighborN(n int) *Value {
	if n > 0 {
		return v.RightNeighborN(n)
	}
	return v.LeftNeighborN(n)
}

func Dump(start *Value) {
	current := start
	for {
		fmt.Printf("%3d", current.number)
		current = current.right
		if current == start {
			break
		}
	}
	fmt.Println("")
}

func part1(data []int) int {
	var start *Value
	var current *Value
	var zeroValue *Value

	order := make([]*Value,len(data))
	for i, originalNumber := range data {
		val := &Value{number: originalNumber}
		order[i] = val
		if originalNumber == 0 {
			zeroValue = val
		}
		if start == nil {
			start = val
			current = val
			continue
		}
		current.right = val
		val.left = current
		current = val
	}

	end := current

	// connect start and end to form a doubly-linked ring-list
	start.left = end
	end.right = start

	for _, val := range order {
			if val.number == 0 {
				continue
			}
			
			target := val.GetNeighborN(val.number)
			val.Move(target, target.right)
	}

	// Dump(start)
	
	n1 := zeroValue.GetNeighborN(1000 % len(data)).number
	n2 := zeroValue.GetNeighborN(2000 % len(data)).number
	n3 := zeroValue.GetNeighborN(3000 % len(data)).number

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	return n1 + n2 + n3
}

func part2(data []string) int {
	return 0
}

func main() {

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(input.LoadInt("input1")))
	fmt.Println(part1(input.LoadInt("input")))
	fmt.Println("     bad: 0, 10909")
	fmt.Println("too high: 12787")

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
