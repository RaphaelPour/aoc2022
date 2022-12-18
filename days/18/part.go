package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
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
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
