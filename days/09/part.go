package main

import (
	"fmt"
	"strings"

	real_math "math"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Point struct {
	x, y int
}

func (p Point) ManhattanDistance(other Point) int {
	return math.Abs(p.x-other.x) + math.Abs(p.y-other.y)
}

func (p Point) EuclideanDistance(other Point) float64 {
	return real_math.Sqrt(float64(math.Pow(math.Abs(p.x-other.x), 2) + math.Pow(math.Abs(p.y-other.y), 2)))
}

func (p Point) Plus(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type Cache map[Point]bool

func (c Cache) Add(item Point) {
	c[item] = true
}

func (c Cache) Count() int {
	return len(c)
}

type Rope struct {
	head, tail Point
	history    Cache
}

func NewRope() *Rope {
	r := new(Rope)
	r.history = make(Cache)
	return r
}

func (r Rope) Dump() {
	for y := -5; y < 5; y++ {
		for x := -5; x < 5; x++ {
			current := Point{x, y}
			if current == r.head && current == r.tail {
				fmt.Print("B")
			} else if current == r.head {
				fmt.Print("H")
			} else if current == r.tail {
				fmt.Print("T")
			} else if _, ok := r.history[current]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (r *Rope) Move(direction string, steps int) {
	for ; steps > 0; steps-- {
		// r.Dump()
		// fmt.Println("Pre  H:", r.head, "T:", r.tail)
		// move head
		switch direction {
		case "R":
			r.head.x += 1
		case "L":
			r.head.x -= 1
		case "U":
			r.head.y -= 1
		case "D":
			r.head.y += 1
		}
		// fmt.Println("Head H:", r.head, "T:", r.tail)

		// early return if head covers the tail or is direct neighbour
		if r.head == r.tail || r.head.EuclideanDistance(r.tail) < 2 {
			continue
		}

		// move tail by minimizing its distance to head
		newTail := r.tail
		distance := 100.0
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				newPoint := r.tail.Plus(Point{x, y})

				// don't cover head, head can only cover tail itself
				if newPoint == r.head {
					continue
				}
				newDist := newPoint.EuclideanDistance(r.head)
				// fmt.Println(newPoint, r.head, newDist)
				if newDist < distance {
					newTail = newPoint
					distance = newDist
				}
			}
		}

		r.history.Add(r.tail)
		r.tail = newTail
		// fmt.Println("Tail H:", r.head, "T:", r.tail)
	}
	// add last element
	r.history.Add(r.tail)
}

func (r *Rope) MoveTail(other Rope) {
	// early return if head covers the tail or is direct neighbour
	if other.tail == r.tail || other.tail.EuclideanDistance(r.tail) < 2 {
		return
	}

	// move tail by minimizing its distance to head
	newTail := r.tail
	distance := 100.0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			newPoint := r.tail.Plus(Point{x, y})

			// don't cover head, head can only cover tail itself
			if newPoint == other.tail {
				continue
			}
			newDist := newPoint.EuclideanDistance(other.tail)
			// fmt.Println(newPoint, r.head, newDist)
			if newDist < distance {
				newTail = newPoint
				distance = newDist
			}
		}
	}

	r.history.Add(r.tail)
	r.history.Add(newTail)
	r.tail = newTail
	// fmt.Println("Tail H:", r.head, "T:", r.tail)
}

type Chain struct {
	ropes []Rope
}

func NewChain(count int) *Chain {
	c := new(Chain)
	c.ropes = make([]Rope, count)
	for i := range c.ropes {
		c.ropes[i] = *NewRope()
	}
	return c
}

func (c *Chain) Move(direction string, steps int) {
	// move one step at a time and propagate movement to the other ropes
	for ; steps > 0; steps-- {
		c.ropes[0].Move(direction, 1)
		for i := 1; i < len(c.ropes); i++ {
			c.ropes[i].MoveTail(c.ropes[i-1])
		}
	}
}

func (c Chain) History() map[Point]bool {
	history := make(map[Point]bool)

	for _, tail := range c.ropes {
		for key := range tail.history {
			history[key] = true
		}
	}
	return history
}

func (c Chain) Count() int {
	return len(c.History())
}

func (c Chain) Dump() {
	hist := c.History()
	horizont := 10
	for y := -horizont; y < horizont; y++ {
		for x := -horizont; x < horizont; x++ {
			current := Point{x, y}
			if _, ok := hist[current]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func part1(data []string) int {
	rope := NewRope()
	for _, line := range data {
		parts := strings.Fields(line)
		direction := parts[0]
		steps := s_strings.ToInt(parts[1])
		rope.Move(direction, steps)
	}
	return rope.history.Count()
}

func part2(data []string) int {
	chain := NewChain(9)
	for _, line := range data {
		parts := strings.Fields(line)
		direction := parts[0]
		steps := s_strings.ToInt(parts[1])
		chain.Move(direction, steps)
		chain.Dump()
	}
	return chain.ropes[8].history.Count()
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
	fmt.Println("too high: 11432")
}
