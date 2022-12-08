package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Grid struct {
	fields [][]int
}

func (g Grid) VisibleCount() int {
	sum := 0
	for y := 0; y < len(g.fields); y++ {
		for x := 0; x < len(g.fields[0]); x++ {
			if g.TreeVisible(x, y) {
				fmt.Printf("%d ", g.fields[y][x])
				sum += 1
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("")
	}
	return sum
}

func (g Grid) TreeVisible(x, y int) bool {
	if !math.Within(y, 0, len(g.fields)-1) || !math.Within(x, 0, len(g.fields[0])-1) {
		panic(fmt.Sprintf("Coordinates out of range: %d/%d", x, y))
	}

	// check if tree is at the border
	/*if x == 0 || x == len(g.fields)-1 || y == 0 || y == len(g.fields[0])-1 {
		return true
	}*/

	// left
	isVisible := true
	for x2 := x - 1; x2 >= 0; x2-- {
		if g.fields[y][x2] >= g.fields[y][x] {
			isVisible = false
		}
	}

	if isVisible {
		return true
	}

	// right
	isVisible = true
	for x2 := x + 1; x2 < len(g.fields[0]); x2++ {
		if g.fields[y][x2] >= g.fields[y][x] {
			isVisible = false
		}
	}

	if isVisible {
		return true
	}

	// top
	isVisible = true
	for y2 := y - 1; y2 >= 0; y2-- {
		if g.fields[y2][x] >= g.fields[y][x] {
			isVisible = false
		}
	}

	if isVisible {
		return true
	}

	// bottom
	isVisible = true
	for y2 := y + 1; y2 < len(g.fields); y2++ {
		if g.fields[y2][x] >= g.fields[y][x] {
			isVisible = false
		}
	}

	return isVisible
}

func (g Grid) ScenicScore() int {
	score := 0
	for y := 0; y < len(g.fields); y++ {
		for x := 0; x < len(g.fields[0]); x++ {
			dist := g.ViewingDistance(x, y)
			if dist > score {
				score = dist
			}
			fmt.Printf("%2d ", dist)
		}
		fmt.Println("")
	}
	return score
}

func (g Grid) ViewingDistance(x, y int) int {
	if !math.Within(y, 0, len(g.fields)-1) || !math.Within(x, 0, len(g.fields[0])-1) {
		panic(fmt.Sprintf("Coordinates out of range: %d/%d", x, y))
	}

	fmt.Println("lel")
	if x == 0 || x == len(g.fields)-1 || y == 0 || y == len(g.fields[0])-1 {
		return 0
	}
	score := 1

	// left
	for x2 := x - 1; x2 >= 0; x2-- {
		if g.fields[y][x2] >= g.fields[y][x] || x2 == 0 {
			score *= math.Abs(x - x2)
			break
		}
	}

	// right
	for x2 := x + 1; x2 < len(g.fields[0]); x2++ {
		if g.fields[y][x2] >= g.fields[y][x] || x2 == len(g.fields[0])-1 {
			score *= math.Abs(x - x2)
			break
		}
	}

	// top
	for y2 := y - 1; y2 >= 0; y2-- {
		if g.fields[y2][x] >= g.fields[y][x] || y2 == 0 {
			score *= math.Abs(y - y2)
			break
		}
	}

	// bottom
	for y2 := y + 1; y2 < len(g.fields); y2++ {
		if g.fields[y2][x] >= g.fields[y][x] || y2 == len(g.fields)-1 {
			score *= math.Abs(y - y2)
			break
		}
	}
	return score
}

func part1(data []string) int {
	grid := Grid{}
	grid.fields = make([][]int, 0)
	for _, line := range data {
		if len(grid.fields) > 0 && len(line) != len(grid.fields[0]) {
			panic(fmt.Sprintf(
				"bad line length: expected %d, got %d",
				len(grid.fields[0]),
				len(line),
			))
		}
		lineFields := make([]int, len(line))
		for i, rawNum := range line {
			num := s_strings.ToInt(string(rawNum))
			lineFields[i] = num
		}
		grid.fields = append(grid.fields, lineFields)
	}

	fmt.Println(grid.fields)
	return grid.VisibleCount()
}

func part2(data []string) int {
	grid := Grid{}
	grid.fields = make([][]int, 0)
	for _, line := range data {
		if len(grid.fields) > 0 && len(line) != len(grid.fields[0]) {
			panic(fmt.Sprintf(
				"bad line length: expected %d, got %d",
				len(grid.fields[0]),
				len(line),
			))
		}
		lineFields := make([]int, len(line))
		for i, rawNum := range line {
			num := s_strings.ToInt(string(rawNum))
			lineFields[i] = num
		}
		grid.fields = append(grid.fields, lineFields)
	}
	grid.ViewingDistance(2, 2)
	return grid.ScenicScore()
	// return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
