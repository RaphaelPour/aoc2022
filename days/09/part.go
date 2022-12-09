package main

import (
	"fmt"
	"strings"

	real_math "math"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	move = map[string]Point{
		"R": Point{1, 0},
		"L": Point{-1, 0},
		"U": Point{0, -1},
		"D": Point{0, 1},
	}
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

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type Rope struct {
	knots   []Point
	history map[Point]bool
}

func NewRope(count int) *Rope {
	r := new(Rope)
	r.knots = make([]Point, count)
	r.history = make(map[Point]bool)

	// rope always starts at 0/0, add this to the history
	r.history[Point{0, 0}] = true
	return r
}

func (r *Rope) Move(direction string, steps int) {
	// move one step at a time and propagate movement to the tail
	for ; steps > 0; steps-- {
		r.moveHead(direction)
		r.moveTail()
	}
}

func (r *Rope) moveHead(direction string) {
	r.knots[0] = r.knots[0].Add(move[direction])
}

func (r *Rope) moveTail() {
	for i := 1; i < len(r.knots); i++ {
		head := r.knots[i-1]
		tail := r.knots[i]

		if head.EuclideanDistance(tail) < 2 {
			continue
		}

		// move tail by minimizing its distance to head
		newTail := tail
		distance := 100.0
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				newPoint := tail.Add(Point{x, y})

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
		r.knots[i] = newTail

		// add new position to history if the tail is the last tail
		if i == len(r.knots)-1 {
			r.history[newTail] = true
		}
	}
}

func (r Rope) Count() int {
	return len(r.history)
}

func part1(data []string) int {
	chain := NewRope(2)
	for _, line := range data {
		parts := strings.Fields(line)
		chain.Move(parts[0], s_strings.ToInt(parts[1]))
	}
	return chain.Count()
}

func part2(data []string) int {
	chain := NewRope(10)
	for _, line := range data {
		parts := strings.Fields(line)
		chain.Move(parts[0], s_strings.ToInt(parts[1]))
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
