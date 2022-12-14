package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

const (
	FREE = iota
	ROCK
	SAND
)

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type Grid struct {
	points      map[Point]int
	abyssHeight int

	min, max Point
}

func NewGrid() Grid {
	g := Grid{}
	g.points = make(map[Point]int)
	return g
}

func (g *Grid) Add(start, end Point) {
	minX := math.Min([]int{start.x, end.x})
	minY := math.Min([]int{start.y, end.y})

	maxX := math.Max([]int{start.x, end.x})
	maxY := math.Max([]int{start.y, end.y})

	fmt.Println(minY)
	if minY < g.abyssHeight {
		g.abyssHeight = minY
	}

	if minX < g.min.x || g.min.x == 0 {
		g.min.x = minX
	}
	if maxX > g.max.x || g.max.x == 0 {
		g.max.x = maxX
	}
	if minY < g.min.y {
		g.min.y = minY
	}
	if maxY > g.max.y {
		g.max.y = maxY
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			g.points[Point{x, y}] = ROCK
		}
	}
}

func (g Grid) IsBlocked(p Point) bool {
	_, ok := g.points[p]
	return ok
}

func (g *Grid) AddSandWithAbyss(p Point) bool {
	// sand is falling to the abyss
	if p.y <= g.abyssHeight {
		return true
	}

	for _, move := range []Point{
		p.Add(Point{0, -1}),  // 1. one step down
		p.Add(Point{-1, -1}), // 2. one step down + left
		p.Add(Point{1, -1}),  // 3. one step down + right
	} {
		// if one move is not blocked -> go in recursion to let sand fall more
		if !g.IsBlocked(move) {
			return g.AddSandWithAbyss(move)
		}
	}
	// sand got blocked
	g.points[p] = SAND
	return false
}

func (g *Grid) AddSandWithGround(p Point) bool {
	if p.y < g.abyssHeight {
		g.points[p] = SAND
		return false
	}
	for _, move := range []Point{
		p.Add(Point{0, -1}),  // 1. one step down
		p.Add(Point{-1, -1}), // 2. one step down + left
		p.Add(Point{1, -1}),  // 3. one step down + right
	} {
		// if one move is not blocked -> go in recursion to let sand fall more
		if !g.IsBlocked(move) {
			return g.AddSandWithGround(move)
		}
	}

	if p.y < g.min.y {
		g.min.y = p.y
	} else if p.y > g.max.y {
		g.max.y = p.y
	}
	if p.x < g.min.x {
		g.min.x = p.x
	} else if p.x > g.max.x {
		g.max.x = p.x
	}

	// sand got blocked
	g.points[p] = SAND

	// if we're at the start, we're done
	if p.x == 500 && p.y == 0 {
		return true
	}

	return false
}

func (g Grid) Dump() {
	fmt.Println(g.min, g.max)
	for y := g.max.y; y >= g.min.y; y-- {
		for x := g.min.x; x <= g.max.x; x++ {
			if val, ok := g.points[Point{x, y}]; ok {
				if val == ROCK {
					fmt.Print("#")
				} else {
					fmt.Print("o")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func part1(data []string) int {
	grid := NewGrid()
	for _, line := range data {
		coords := strings.Split(line, "->")

		for i := 1; i < len(coords); i++ {
			rawStart := strings.Split(strings.TrimSpace(coords[i-1]), ",")
			rawEnd := strings.Split(strings.TrimSpace(coords[i]), ",")

			grid.Add(
				Point{s_strings.ToInt(rawStart[0]), -s_strings.ToInt(rawStart[1])},
				Point{s_strings.ToInt(rawEnd[0]), -s_strings.ToInt(rawEnd[1])},
			)
		}
	}

	for units := 0; ; units++ {
		if grid.AddSandWithAbyss(Point{500, 0}) {
			return units
		}
	}
	return 0
}

func part2(data []string) int {
	grid := NewGrid()
	for _, line := range data {
		coords := strings.Split(line, "->")

		for i := 1; i < len(coords); i++ {
			rawStart := strings.Split(strings.TrimSpace(coords[i-1]), ",")
			rawEnd := strings.Split(strings.TrimSpace(coords[i]), ",")

			grid.Add(
				Point{s_strings.ToInt(rawStart[0]), -s_strings.ToInt(rawStart[1])},
				Point{s_strings.ToInt(rawEnd[0]), -s_strings.ToInt(rawEnd[1])},
			)
		}
	}

	for units := 1; ; units++ {
		if grid.AddSandWithGround(Point{500, 0}) {
			return units
		}
	}
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
