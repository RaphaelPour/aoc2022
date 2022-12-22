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

type Ring struct {
	zero *Value
	values []*Value
}

func NewRing(data []int) Ring {
	var start *Value
	var current *Value

	r := Ring{}
	r.values = make([]*Value,len(data))
	for i, originalNumber := range data {
		val := &Value{number: originalNumber}
		r.values[i] = val
		if originalNumber == 0 {
			r.zero = val
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

	return r
}

func (r *Ring) Mix() {
	for _, val := range r.values {
			if val.number == 0 {
				continue
			}
			
			target := val.GetNeighborN(val.number)

			if val.number > 0 {
				val.Move(target, target.right)
			} else {
				val.Move(target.left, target)
			}
	}
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

	for ; n > 0; n-- {
		current = current.right
	}
	return current
}

func (v Value) LeftNeighborN(n int) *Value {
	current := &v

	for ; n > 0; n-- {
		current = current.left
	}
	return current
}

func (v Value) GetNeighborN(n int) *Value {
	if n == 0 {
		return &v
	} else if  n > 0 {
		return v.RightNeighborN(n)
	}
	return v.LeftNeighborN(-n)
}

func Dump(start *Value) {
	current := start
	for {
		fmt.Printf("%3d,", current.number)
		current = current.right
		if current == start {
			break
		}
	}
	fmt.Println("")
}

func part1(data []int) int {
	ring := NewRing(data)
	ring.Mix()

	n1 := ring.zero.GetNeighborN(1000 % len(data)).number
	n3 := ring.zero.GetNeighborN(3000 % len(data)).number
	n2 := ring.zero.GetNeighborN(2000 % len(data)).number
	
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
	fmt.Println("     bad: 0,646, 963, 10909")
	fmt.Println("too high: 12787")

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
