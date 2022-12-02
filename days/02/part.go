package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

const (
	LOOSE = "X"
	DRAW  = "Y"
	WIN   = "Z"
)

var (
	win = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	draw = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	loose = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	points = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
)

func part1(data []string) int {
	score := 0
	for _, line := range data {
		round := strings.Split(line, " ")
		if win[round[0]] == round[1] {
			score += 6
		} else if draw[round[0]] == round[1] {
			score += 3
		}
		score += points[round[1]]
	}
	return score
}

func part2(data []string) int {
	score := 0
	for _, line := range data {
		round := strings.Split(line, " ")
		turn := ""
		if round[1] == WIN {
			turn = win[round[0]]
			score += 6
		} else if round[1] == DRAW {
			turn = draw[round[0]]
			score += 3
		} else {
			turn = loose[round[0]]
		}
		score += points[turn]
	}
	return score
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
