package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	directions = []Point{
		Point{1, 0, 0},
		Point{-1, 0, 0},
		Point{0, 1, 0},
		Point{0, -1, 0},
		Point{0, 0, 1},
		Point{0, 0, -1},
	}
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

func (p Point) OutOfBounds(min, max Point) bool {
	return p.x <= min.x || p.x >= max.x ||
		p.y <= min.y || p.y >= max.y ||
		p.z <= min.z || p.z >= max.z
}

func (p Point) PossibleNeighbors() []Point {
	neighbors := make([]Point, len(directions))

	for i, newPoint := range directions {
		neighbors[i] = p.Add(newPoint)
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
	max := Point{-10, -10, -10}
	min := Point{10, 10, 10}

	for _, line := range data {
		p := FromLine(line)
		m[p] = struct{}{}
		max = p.Max(max)
		min = p.Min(min)
	}

	min = min.Add(Point{-1, -1, -1})
	max = max.Add(Point{1, 1, 1})

	airCubes := make(map[Point]bool)
	lavaCubes := make(map[Point]struct{})
	airCubes[min] = false

	change := true
	sides := 0
	for change {
		change = false

		for cube, processed := range airCubes {
			if processed {
				continue
			}

			change = true

			for _, neighbor := range cube.PossibleNeighbors() {
				// skip neighbors outside of bounds
				if max.Max(neighbor) != max || min.Min(neighbor) != min {
					continue
				}

				// skip already found air cubes
				if _, ok := airCubes[neighbor]; ok {
					continue
				}

				// add side and continue
				if _, ok := lavaCubes[neighbor]; ok {
					sides++
					continue
				}

				if _, ok := m[neighbor]; ok {
					lavaCubes[neighbor] = struct{}{}
					sides++
				} else {
					airCubes[neighbor] = false
				}
			}
			airCubes[cube] = true
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
	fmt.Println("bad: 2571, 2987")
	fmt.Println("too high: 3719")
}
