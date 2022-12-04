package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Range struct {
	from, to int
}

func NewRange(rawRange string) (*Range, error) {
	rawParts := strings.Split(rawRange, "-")
	if len(rawParts) != 2 {
		return nil, fmt.Errorf("error parsing '%s': must have two numbers split with -", rawRange)
	}

	return &Range{s_strings.ToInt(rawParts[0]), s_strings.ToInt(rawParts[1])}, nil
}

func (r1 Range) Contains(r2 Range) bool {
	if r1.from <= r2.from && r1.to >= r2.to {
		return true
	}

	if r2.from <= r1.from && r2.to >= r1.to {
		return true
	}
	return false
}

func (r1 Range) Overlap(r2 Range) bool {
	if r1.Contains(r2) {
		return true
	}

	if r1.from <= r2.from && r1.to >= r2.to {
		return true
	}

	if r1.from <= r2.from && r1.to >= r2.from {
		return true
	}

	if r2.from <= r1.from && r2.to >= r1.from {
		return true
	}

	if r2.from <= r1.from && r2.to >= r1.to {
		return true
	}

	if r1.to == r2.from || r2.to == r1.from {
		return true
	}

	return false
}

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			panic(fmt.Sprintf("pair should have two elements, got %d ('%s')", len(pair), pair))
		}

		r1, err := NewRange(pair[0])
		if err != nil {
			panic(fmt.Sprintf("error parsing first pair %s: %s", pair, err))
		}

		r2, err := NewRange(pair[1])
		if err != nil {
			panic(fmt.Sprintf("error parsing second pair %s: %s", pair, err))
		}

		if r1.Contains(*r2) {
			sum += 1
		}
	}
	return sum
}

func part2(data []string) int {
	sum := 0
	for _, line := range data {
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			panic(fmt.Sprintf("pair should have two elements, got %d ('%s')", len(pair), pair))
		}

		r1, err := NewRange(pair[0])
		if err != nil {
			panic(fmt.Sprintf("error parsing first pair %s: %s", pair, err))
		}

		r2, err := NewRange(pair[1])
		if err != nil {
			panic(fmt.Sprintf("error parsing second pair %s: %s", pair, err))
		}

		if r1.Overlap(*r2) {
			sum += 1
		}
	}
	return sum
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println("bad: 544")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
