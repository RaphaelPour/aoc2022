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
	knots []Point
	history    Cache
}

func NewChain(count int) *Chain {
	c := new(Chain)
	c.knots = make([]Point, count)
	c.history = make(Cache)
	return c
}

func (c *Chain) Move(direction string, steps int) {
	// move one step at a time and propagate movement to the other ropes
	for ; steps > 0; steps-- {
		c.MoveHead(direction)
		c.MoveTail()
	}
}

func (c *Chain) MoveHead(direction string) {
		switch direction {
		case "R":
			c.knots[0].x += 1
		case "L":
			c.knots[0].x -= 1
		case "U":
			c.knots[0].y -= 1
		case "D":
			c.knots[0].y += 1
		}
}

func (c *Chain) MoveTail() {
	for i:=1;i<len(c.knots);i++ {
		head := c.knots[i-1]
		tail := c.knots[i]

		// add current position to history if the tail is the last tail
		if i == len(c.knots)-1 {
			c.history.Add(tail)
		}

		if head == tail || head.EuclideanDistance(tail) < 2 {
			continue
		}

		// move tail by minimizing its distance to head
		newTail := tail
		distance := 100.0
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				newPoint := tail.Plus(Point{x, y})

				// don't cover head, head can only cover tail itself
				if newPoint == head {
					continue
				}
				newDist := newPoint.EuclideanDistance(head)
				// fmt.Println(newPoint, r.head, newDist)
				if newDist < distance {
					newTail = newPoint
					distance = newDist
				}
			}
		}
		c.knots[i] = newTail

		// add new position to history if the tail is the last tail
		if i == len(c.knots)-1 {
			c.history.Add(newTail)
		}
	}
}

func (c Chain) Count() int {
	return len(c.history)
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
	chain := NewChain(10)
	for _, line := range data {
		parts := strings.Fields(line)
		direction := parts[0]
		steps := s_strings.ToInt(parts[1])
		chain.Move(direction, steps)
	}
	return chain.Count()
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
