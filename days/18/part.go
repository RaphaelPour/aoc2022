package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Point struct {
	x, y, z int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	p.z += other.z
	return p
}

func (p Point) Max(other Point) Point {
	p.x = math.Max([]int{p.x, other.x})
	p.y = math.Max([]int{p.y, other.y})
	p.z = math.Max([]int{p.z, other.z})
	return p
}

func (p Point) Min(other Point) Point {
	p.x = math.Min([]int{p.x, other.x})
	p.y = math.Min([]int{p.y, other.y})
	p.z = math.Min([]int{p.z, other.z})
	return p
}

func (p Point) PossibleNeighbors() []Point {
	neighbors := make([]Point, 0)

	for _, newPoint := range []Point{
		p.Add(Point{1, 0, 0}),
		p.Add(Point{-1, 0, 0}),
		p.Add(Point{0, 1, 0}),
		p.Add(Point{0, -1, 0}),
		p.Add(Point{0, 0, 1}),
		p.Add(Point{0, 0, -1}),
	} {
		neighbors = append(neighbors, newPoint)
	}

	return neighbors
}

func FromLine(line string) Point {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		panic(fmt.Sprintf("expected line '%s' to have three components, got %d", line, len(parts)))
	}

	return Point{
		s_strings.ToInt(parts[0]),
		s_strings.ToInt(parts[1]),
		s_strings.ToInt(parts[2]),
	}
}

func part1(data []string) int {
	m := make(map[Point]struct{})

	// parse all cubes
	for _, line := range data {
		m[FromLine(line)] = struct{}{}
	}

	sides := 0
	for cube := range m {
		for _, neighbor := range cube.PossibleNeighbors() {
			if _, ok := m[neighbor]; !ok {
				sides++
			}
		}
	}

	return sides
}

func part2(data []string) int {

	// parse all cubes
	m := make(map[Point]struct{})
	min, max := Point{10, 10, 10}, Point{-10, -10, -10}
	for _, line := range data {
		m[FromLine(line)] = struct{}{}

	}

	sides := 0
	maybePocket := make(map[Point]int)
	for cube := range m {
		for _, neighbor := range cube.PossibleNeighbors() {
			if _, ok := m[neighbor]; !ok {
				sides++
				if _, ok := maybePocket[neighbor]; !ok {
					maybePocket[neighbor] = 0
				}
				maybePocket[neighbor]++
			}
		}
	}

	// subtract air pockets
	for pocket, count := range maybePocket {
		// process 1x1x1 air pockets
		if count == 6 {
			sides -= 6
			continue
		}

		trapped := true
		realSides := 0
		for _, neighbor := range pocket.PossibleNeighbors() {
			_, ok1 := maybePocket[neighbor]
			_, ok2 := m[neighbor]

			if !(ok1 || ok2) {
				trapped = false
				break
			}

			if ok2 {
				realSides++
			}
		}
		if trapped {
			sides -= realSides
		}
	}

	return sides
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
	fmt.Println("bad: 2987")
	fmt.Println("too high: 3719")
}
