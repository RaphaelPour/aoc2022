package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

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
			fmt.Printf("%2d",height)
		}
		fmt.Println("")
	}

	// breadth first search
	path := make(map[Point]bool)
	queue := make([]Point,0)
	queue = append(queue, h.start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == h.goal {
			fmt.Println("GOAL")
			path[current] = true
			Dump(path)
			return len(path)
		}

		for _, neighbor := range []Point{
			current.Add(Point{-1,0}),
			current.Add(Point{1,0}),
			current.Add(Point{0,-1}),
			current.Add(Point{0,1}),
		}{
			if h.IsOutOfBounds(neighbor) {
				continue
			}
			
			if h.Get(neighbor) - h.Get(current) > 1 {
				continue
			}

			if _, visited := path[neighbor]; visited {
				continue
			}

			path[current] = true
			queue = append(queue, neighbor)
		}
	}
	return 0
}

func main() {
	data := input.LoadString("input1")
	fmt.Println(part1(data))
}
