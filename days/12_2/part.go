package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Field struct  {
	p Point
	dist int
}

type Point struct {
	x,y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type HeightMap struct {
	grid [][]int
	start, goal Point
}

func NewHeightMap(rows int) HeightMap {
	h := HeightMap{}
	h.grid = make([][]int, rows)
	return h
}

func (h HeightMap) Get(p Point) int {
	return h.grid[p.y][p.x]
}

func (h HeightMap) IsOutOfBounds(p Point) bool {
	return p.x < 0 || p.x >= len(h.grid[0]) || p.y < 0 || p.y >= len(h.grid)
}

func Dump(path map[Point]bool) {
	for y := 0;y<5;y++ {
		for x := 0 ;x<8;x++ {
			if _, ok := path[Point{x,y,}]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (h HeightMap) search(start Point) int {
	path := make(map[Point]bool)
	queue := make([]Field,0)
	queue = append(queue, Field{start,0})
	i := 0
	for len(queue) > 0 {
		i++
		current := queue[0]
		queue = queue[1:]
		if current.p == h.goal {
			path[current.p] = true
			return current.dist
		}

		for _, neighbor := range []Point{
			current.p.Add(Point{-1,0}),
			current.p.Add(Point{1,0}),
			current.p.Add(Point{0,-1}),
			current.p.Add(Point{0,1}),
		}{
			if h.IsOutOfBounds(neighbor) {
				continue
			}
			
			if h.Get(neighbor) - h.Get(current.p) > 1 {
				continue
			}

			if _, visited := path[neighbor]; visited {
				continue
			}

			path[neighbor] = true
			queue = append(queue, Field{neighbor, current.dist+1})
		}
	}

	return 0
}

func part1(data []string) int  {
	h := NewHeightMap(len(data))

	// parse input
	for y, line := range data {
		h.grid[y] = make([]int, len(line))
		for x, field := range line {
			height := int(field - 'a')
			if field == 'S' {
				height = 0
				h.start = Point{x,y}
			} else if field == 'E' {
				height = int('z' - 'a')
				h.goal = Point{x,y}
			}
			h.grid[y][x] = height
		}
	}

	return h.search(h.start)
}

func part2(data []string) int {
	h := NewHeightMap(len(data))

	// parse input
	starts := make([]Point,0)
	for y, line := range data {
		h.grid[y] = make([]int, len(line))
		for x, field := range line {
			height := int(field - 'a')
			if field == 'S' {
				height = 0
				h.start = Point{x,y}
			} else if field == 'E' {
				height = int('z' - 'a')
				h.goal = Point{x,y}
			}
			if height == 0{
				starts = append(starts, Point{x,y})
			}
			h.grid[y][x] = height
		}
	}

	minSteps := -1
	for _, start := range starts {
		steps := h.search(start)
		if steps > 0 && (minSteps == -1 || steps < minSteps) {
			minSteps = steps
		}
	}
	return minSteps
}

func main() {
	data := input.LoadString("input")
	fmt.Println("== Part 1 ==")
	fmt.Println(part1(data))

	fmt.Println("== Part 2 ==")
	fmt.Println(part2(data))
}
