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

func (v Value) GetNeighborN(n int) *Value {
	current := &v
	for i := 0; i < n; i++ {
		current = current.right
	}
	return current
}

func Dump(values []*Value) {
	for _, val := range values {
		fmt.Println(val)
	}
}

func part1(data []int) int {
	var start *Value
	var current *Value
	values := make([]*Value, len(data))
	for i, originalNumber := range data {
		val := &Value{number: originalNumber}
		values[i] = val
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
	Dump(values)
	fmt.Println("parsing [x]")

	allProcessed := false
	for !allProcessed {
		allProcessed = true
		for _, current := range values {
			if current.moved {
				continue
			}
			allProcessed = false

			to := current.number - 1
			if to < 0 {
				to = len(data) + to - 1
			}

			to = to % len(data)

			target := current.GetNeighborN(to)
			current.Move(target, target.right)
			current.moved = true
		}
	}

	// Dump(values)

	var zeroValue *Value
	current, start = values[0], values[0]
	for {
		// fmt.Println(current)
		if current.number == 0 {
			zeroValue = current
			break
		}

		current = current.right
		if current == start {
			break
		}
	}

	n1 := zeroValue.GetNeighborN(1000 % len(values))
	n2 := zeroValue.GetNeighborN(2000 % len(values))
	n3 := zeroValue.GetNeighborN(3000 % len(values))

	fmt.Println("n1:", n1.number)
	fmt.Println("n2:", n2.number)
	fmt.Println("n3:", n3.number)

	return n1.number + n2.number + n3.number
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadInt("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
