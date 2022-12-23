package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

const (
	NORTH = iota
	SOUTH
	WEST
	EAST
	DIRECTION_COUNT
)

var (
	northMove = Point{0, -1}
	southMove = Point{0, 1}
	westMove  = Point{-1, 0}
	eastMove  = Point{1, 0}

	northEastMove = northMove.Add(eastMove)
	northWestMove = northMove.Add(westMove)

	southEastMove = southMove.Add(eastMove)
	southWestMove = southMove.Add(westMove)

	moves = map[int][]Point{
		NORTH: []Point{northMove, northEastMove, northWestMove},
		SOUTH: []Point{southMove, southEastMove, southWestMove},
		WEST:  []Point{westMove, northWestMove, southWestMove},
		EAST:  []Point{eastMove, northEastMove, southEastMove},
	}

	dir2str = []string{"N", "S", "W", "E"}
)

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

func (p Point) Max(other Point) Point {
	if other.x > p.x {
		p.x = other.x
	}

	if other.y > p.y {
		p.y = other.y
	}
	return p
}

func (p Point) Min(other Point) Point {
	if other.x < p.x {
		p.x = other.x
	}

	if other.y < p.y {
		p.y = other.y
	}
	return p
}

func (p Point) Area() int {
	return p.x * p.y
}

type Field struct {
	currentDirection int
}

func (f *Field) NextDirection() {
	oldDir := dir2str[f.currentDirection]
	f.currentDirection = (f.currentDirection + 1) % DIRECTION_COUNT
	fmt.Println("%s -> %s\n", oldDir, f.currentDirection)
}

type Board struct {
	fields map[Point]Field

	currentDirection int
	min, max         Point
}

func NewBoard(data []string) Board {
	b := Board{}
	b.fields = make(map[Point]Field)
	b.min = Point{1, 1}
	b.max = Point{0, 0}
	b.currentDirection = NORTH
	for y, row := range data {
		for x, cell := range row {
			b.min = b.min.Min(Point{x, y})
			b.max = b.max.Max(Point{x, y})
			if string(cell) != "#" {
				continue
			}
			b.fields[Point{x, y}] = Field{NORTH}
		}
	}

	return b
}

func (b Board) Dimensions() Point {
	min, max := Point{10, 10}, Point{0, 0}

	for p := range b.fields {
		min = min.Min(p)
		max = max.Max(p)
	}
	return Point{math.Abs(max.x - min.x), math.Abs(max.y - min.y)}
}

func (b *Board) ProposeMove(p Point, f *Field) (Point, bool) {
	for dir := b.currentDirection; dir < b.currentDirection+DIRECTION_COUNT; dir++ {
		free := true
		for _, move := range moves[dir%DIRECTION_COUNT] {
			newPoint := p.Add(move)
			if _, found := b.fields[newPoint]; found {
				free = false
				break
			}
		}
		if free {
			return p.Add(moves[dir%DIRECTION_COUNT][0]), true
		}
	}
	return p, false
}

func (b *Board) Next() {
	fmt.Println("move", dir2str[b.currentDirection])
	buffer := make(map[Point][]Point)
	for p, field := range b.fields {
		newPoint, _ := b.ProposeMove(p, &field)

		if _, ok := buffer[newPoint]; !ok {
			buffer[newPoint] = make([]Point, 0)
		}
		buffer[newPoint] = append(buffer[newPoint], p)
	}

	for p, oldPoints := range buffer {
		// skip positions that have multiple fields
		if len(oldPoints) > 1 {
			continue
		}

		b.min = b.min.Min(p)
		b.max = b.max.Max(p)

		b.fields[p] = b.fields[oldPoints[0]]
		delete(b.fields, oldPoints[0])
	}

	b.currentDirection = (b.currentDirection + 1) % DIRECTION_COUNT
}

func (b Board) Dump() {
	min, max := Point{10, 10}, Point{0, 0}

	for p := range b.fields {
		min = min.Min(p)
		max = max.Max(p)
	}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if _, ok := b.fields[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func part1(data []string) int {
	board := NewBoard(data)

	board.Dump()
	for round := 1; round <= 5; round++ {
		fmt.Printf(" == Round %d ==\n", round)
		board.Next()
		board.Dump()
	}

	return board.Dimensions().Area() - len(board.fields)
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input2")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
