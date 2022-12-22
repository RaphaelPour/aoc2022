package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	cursor = map[int]string{
		RIGHT: ">",
		DOWN:  "v",
		LEFT:  "<",
		UP:    "^",
	}
)

const (
	VOID = iota
	BLOCK
	FREE
)

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
	DIRECTION_BOUNDARY
)

const (
	MOVE = iota
	TURN
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d/%d)", p.x, p.y)
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

func (p Point) Max(other Point) Point {
	p.x = math.Max[int]([]int{p.x, other.x})
	p.y = math.Max[int]([]int{p.y, other.y})
	return p
}

func (p Point) Min(other Point) Point {
	p.x = math.Min[int]([]int{p.x, other.x})
	p.y = math.Min[int]([]int{p.y, other.y})
	return p
}

func (p Point) Move(direction int) Point {
	switch direction {
	case LEFT:
		p.x--
	case RIGHT:
		p.x++
	case UP:
		p.y--
	case DOWN:
		p.y++
	}
	return p
}

type Move struct {
	kind  int
	steps int
	turn  int
}

func (m Move) String() string {
	if m.kind == TURN {
		return fmt.Sprintf("turn %s", cursor[m.turn])
	}

	return fmt.Sprintf("move %d steps", m.steps)
}

type Board struct {
	start  Point
	fields map[Point]int
	moves  []Move

	past     map[Point]string
	min, max Point
}

func NewBoard(data []string) Board {
	var endOfBlock int
	startInit := false
	b := Board{}
	b.min = Point{100, 100}
	b.max = Point{0, 0}
	b.fields = make(map[Point]int)
	b.past = make(map[Point]string)
	for y, row := range data {
		if row == "" {
			endOfBlock = y
			break
		}
		for x, rawCell := range row {
			var fieldType int
			switch string(rawCell) {
			case " ":
				continue
			case "#":
				fieldType = BLOCK
			case ".":
				b.min = b.min.Min(Point{x + 1, y + 1})
				b.max = b.max.Max(Point{x + 1, y + 1})
				fieldType = FREE
				if !startInit {
					b.start = Point{x + 1, y + 1}
					startInit = true
				}
			default:
				panic(fmt.Sprintf("unknown field type %s", string(rawCell)))
			}

			b.fields[Point{x + 1, y + 1}] = fieldType
		}
	}

	pattern := regexp.MustCompile(`(\d+|[RL])`)
	match := pattern.FindAllStringSubmatch(data[endOfBlock+1], -1)
	b.moves = make([]Move, len(match))
	for i, m := range match {
		switch m[1] {
		case "R":
			b.moves[i] = Move{kind: TURN, turn: RIGHT}
		case "L":
			b.moves[i] = Move{kind: TURN, turn: LEFT}
		default:
			b.moves[i] = Move{kind: MOVE, steps: s_strings.ToInt(m[1])}
		}
	}

	return b
}

func (b Board) Wrap(p Point, direction int) (Point, bool) {
	switch direction {
	case RIGHT:
		for x := b.min.x; x < p.x; x++ {
			if kind, ok := b.fields[Point{x, p.y}]; ok && kind != VOID {
				if kind == BLOCK {
					return p, false
				}
				return Point{x, p.y}, true
			}
		}
	case LEFT:
		for x := b.max.x; x > p.x; x-- {
			if kind, ok := b.fields[Point{x, p.y}]; ok && kind != VOID {
				if kind == BLOCK {
					return p, false
				}
				return Point{x, p.y}, true
			}
		}
	case DOWN:
		for y := b.min.y; y < p.y; y++ {
			if kind, ok := b.fields[Point{p.x, y}]; ok && kind != VOID {
				if kind == BLOCK {
					return p, false
				}
				return Point{p.x, y}, true
			}
		}
	case UP:
		for y := b.max.y; y > p.y; y-- {
			if kind, ok := b.fields[Point{p.x, y}]; ok && kind != VOID {
				if kind == BLOCK {
					return p, false
				}
				return Point{p.x, y}, true
			}
		}
	}
	return p, false
}
func (b *Board) Move() (Point, int) {
	currentPosition := b.start
	currentDirection := RIGHT
	b.past[currentPosition] = cursor[currentDirection]

	for _, move := range b.moves {
		//b.Dump()
		// fmt.Println(move, currentPosition, cursor[currentDirection])
		if move.kind == TURN {
			if move.turn == RIGHT {
				currentDirection = (currentDirection + 1) % DIRECTION_BOUNDARY
			} else {
				currentDirection = (currentDirection + 3) % DIRECTION_BOUNDARY
			}
			b.past[currentPosition] = cursor[currentDirection]
			continue
		}

		for step := 0; step < move.steps; step++ {
			newPosition := currentPosition.Move(currentDirection)
			field, found := b.fields[newPosition]

			if found && field == BLOCK {
				break
			}

			if found && field == FREE {
				currentPosition = newPosition
				b.past[currentPosition] = cursor[currentDirection]
				// fmt.Println(move, currentPosition, cursor[currentDirection])
				continue
			}

			//  Otherwise we need to wrap around
			// fmt.Println("WRAP")
			wrapped, ok := b.Wrap(newPosition, currentDirection)
			if !ok {
				continue
			}
			currentPosition = wrapped
			b.past[currentPosition] = cursor[currentDirection]
		}
	}

	fmt.Println("final:", currentPosition, currentDirection)
	return currentPosition, currentDirection
}

func (b Board) Dump() {
	for y := b.min.y; y <= b.max.y; y++ {
		for x := b.min.x; x <= b.max.x; x++ {
			if field, ok := b.past[Point{x, y}]; ok {
				fmt.Print(field)
				continue
			} else if field, ok := b.fields[Point{x, y}]; ok {
				switch field {
				case VOID:
					fmt.Print(" ")
				case FREE:
					fmt.Print(".")
				case BLOCK:
					fmt.Print("#")
				}
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func part1(data []string) int {
	board := NewBoard(data)
	p, d := board.Move()
	board.Dump()
	return 1000*p.y + 4*p.x + d
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))
	fmt.Println("too high: 45496")

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
