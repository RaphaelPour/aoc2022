package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/queue"
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

type Field struct {
	p      Point
	parent *Field
	dist   int
}

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type HeightMap struct {
	grid        [][]int
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
	for y := 0; y < 5; y++ {
		for x := 0; x < 8; x++ {
			if _, ok := path[Point{x, y}]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (h HeightMap) search(start Point) (int, map[Point]struct{}) {
	path := make(map[Point]bool)
	queue := queue.NewQueue[Field]()
	queue.Enqueue(Field{start, nil, 0})
	i := 0
	for len(queue) > 0 {
		i++
		current := queue.Dequeue()
		if current.p == h.goal {
			path[current.p] = true
			goalPath := make(map[Point]struct{})
			iterator := &current
			for iterator.parent != nil {
				goalPath[iterator.p] = struct{}{}
				iterator = iterator.parent
			}
			goalPath[iterator.p] = struct{}{}
			return current.dist, goalPath
		}

		for _, neighbor := range []Point{
			current.p.Add(Point{-1, 0}),
			current.p.Add(Point{1, 0}),
			current.p.Add(Point{0, -1}),
			current.p.Add(Point{0, 1}),
		} {
			if h.IsOutOfBounds(neighbor) {
				continue
			}

			if h.Get(neighbor)-h.Get(current.p) > 1 {
				continue
			}

			if _, visited := path[neighbor]; visited {
				continue
			}

			path[neighbor] = true
			queue.Enqueue(Field{neighbor, &current, current.dist + 1})
		}
	}

	return 0, nil
}

func part1(data []string) int {
	h := NewHeightMap(len(data))

	// parse input
	for y, line := range data {
		h.grid[y] = make([]int, len(line))
		for x, field := range line {
			height := int(field - 'a')
			if field == 'S' {
				height = 0
				h.start = Point{x, y}
			} else if field == 'E' {
				height = int('z' - 'a')
				h.goal = Point{x, y}
			}
			h.grid[y][x] = height
		}
	}

	dist, _ := h.search(h.start)
	return dist
}

func part2(data []string) (int, map[Point]struct{}) {
	h := NewHeightMap(len(data))

	// parse input
	starts := make([]Point, 0)
	for y, line := range data {
		h.grid[y] = make([]int, len(line))
		for x, field := range line {
			height := int(field - 'a')
			if field == 'S' {
				height = 0
				h.start = Point{x, y}
			} else if field == 'E' {
				height = int('z' - 'a')
				h.goal = Point{x, y}
			}
			if height == 0 {
				starts = append(starts, Point{x, y})
			}
			h.grid[y][x] = height
		}
	}

	minSteps := -1
	var bestPath map[Point]struct{}
	for _, start := range starts {
		steps, path := h.search(start)
		if steps > 0 && (minSteps == -1 || steps < minSteps) {
			minSteps = steps
			bestPath = path
		}
	}
	return minSteps, bestPath
}

func renderSTL(input []string) {
	fmt.Println("width:", len(input[0]))
	fmt.Println("height:", len(input))
	filename := "day12.stl"
	_, path := part2(input)
	boxes := make([]sdf.SDF3, 0)

	plate2d := sdf.Box2D(v2.Vec{float64(len(input[0])), float64(len(input))}, 1)
	plate3d := sdf.Extrude3D(plate2d, 1.0)
	plateM := sdf.Translate3d(v3.Vec{
		float64(len(input[0])) / 2,
		float64(len(input)) / 2,
		0,
	})

	boxes = append(boxes, sdf.Transform3D(plate3d, plateM))
	for y, row := range input {
		for x, rawCell := range row {
			cell := float64(rawCell - 'a')
			if cell == 0 {
				continue
			}

			if _, ok := path[Point{x, y}]; ok {
				cell -= 1.0
			}

			box2d := sdf.Box2D(v2.Vec{0.5, 0.5}, 0)
			// add one so level 0 has one unit
			height := cell + 0.5
			box3d := sdf.Extrude3D(box2d, height)
			boxes = append(boxes, sdf.Transform3D(box3d,
				sdf.Translate3d(v3.Vec{float64(x), float64(y), height / 2}),
			))
			boxes = append(boxes, sdf.Transform3D(box3d,
				sdf.Translate3d(v3.Vec{float64(x) + 0.5, float64(y), height / 2}),
			))
			boxes = append(boxes, sdf.Transform3D(box3d,
				sdf.Translate3d(v3.Vec{float64(x), float64(y) + 0.5, height / 2}),
			))
			boxes = append(boxes, sdf.Transform3D(box3d,
				sdf.Translate3d(v3.Vec{float64(x) + 0.5, float64(y) + 0.5, height / 2}),
			))

		}
	}

	fmt.Printf("generated %d boxes\n", len(boxes))
	start := time.Now()
	render.ToSTL(sdf.Union3D(boxes...), filename, render.NewMarchingCubesUniform(100))
	fmt.Printf("needed %s\n", time.Since(start))
}

func renderHeightMap(input []string) {
	fmt.Println("width:", len(input[0]))
	fmt.Println("height:", len(input))
	_, path := part2(input)
	image := image.NewNRGBA(image.Rect(0, 0, len(input[0]), len(input)))
	for y, row := range input {
		for x, rawCell := range row {
			cell := int(rawCell - 'a')
			// 0 should be white and 9 very black
			// spread the gray values accors the whole range
			c := uint8(255.0 - (255.0 / float64(cell)))
			clr := color.NRGBA{R: c, G: c, B: c, A: 255}
			if _, ok := path[Point{x, y}]; ok {
				clr = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			}
			image.Set(x, y, clr)
		}
	}

	filename := fmt.Sprintf("day12_%d.png", time.Now().Unix())
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, image); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	data := input.LoadString("input")
	fmt.Println("== Part 1 ==")
	fmt.Println(part1(data))

	fmt.Println("== Part 2 ==")
	steps, _ := part2(data)
	fmt.Println(steps)
	// renderHeightMap(data)
	renderSTL(data)
}
