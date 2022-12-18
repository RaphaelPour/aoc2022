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

type Neighbor struct {
	point   Point
	leaking bool
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

func (p Point) PossibleNeighbors() []Neighbor {
	neighbors := make([]Neighbor, 0)

	for _, newPoint := range directions {
		neighbors = append(neighbors, Neighbor{point: newPoint})
	}

	return neighbors
}

func (p Point) Terminates(pockets map[Point]int, cubes map[Point]struct{}) bool {
	for _, dir := range directions {
		if !p.TerminatesInDirection(dir, pockets, cubes) {
			return false
		}
	}
	return true
}

func (p Point) TerminatesInDirection(direction Point, pockets map[Point]int, cubes map[Point]struct{}) bool {
	return false
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
			if _, ok := m[neighbor.point]; !ok {
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
		p := FromLine(line)
		m[p] = struct{}{}
		min = min.Min(p)
		max = max.Max(p)
	}

	sides := 0
	maybePocket := make(map[Neighbor]int)
	for cube := range m {
		for _, neighbor := range cube.PossibleNeighbors() {
			if _, ok := m[neighbor.point]; !ok {
				sides++
				if _, ok := maybePocket[neighbor]; !ok {
					maybePocket[neighbor] = 0
				}
				maybePocket[neighbor]++
			}
		}
	}

	change := true
	for change {
		change = false
		// subtract air pockets
		for pocket, count := range maybePocket {
			// process 1x1x1 air pockets
			if count == 6 {
				sides -= 6
				continue
			}

			for _, neighbor := range pocket.point.PossibleNeighbors() {
				_, ok := maybePocket[neighbor]
				if neighbor.point.OutOfBounds(min, max) || (ok && neighborPocket.leaking) {
					pocket.leaking = true
					change = true
					break
				}
			}
		}
	}

	for _, pocket := range maybePocket {
		if !pocket.leaking {
			sides--
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
