package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/span"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			panic(fmt.Sprintf("pair should have two elements, got %d ('%s')", len(pair), pair))
		}

		rawParts := strings.Split(pair[0], "-")
		if len(rawParts) != 2 {
			panic(fmt.Errorf("error parsing first span '%s': must have two numbers split with -", pair[2]))
		}
		s1 := &span.Span{s_strings.ToInt(rawParts[0]), s_strings.ToInt(rawParts[1])}

		rawParts = strings.Split(pair[1], "-")
		if len(rawParts) != 2 {
			panic(fmt.Errorf("error parsing second span '%s': must have two numbers split with -", pair[1]))
		}
		s2 := &span.Span{s_strings.ToInt(rawParts[0]), s_strings.ToInt(rawParts[1])}

		if s1.Contains(*s2) {
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

		rawParts := strings.Split(pair[0], "-")
		if len(rawParts) != 2 {
			panic(fmt.Errorf("error parsing first span '%s': must have two numbers split with -", pair[2]))
		}
		s1 := &span.Span{s_strings.ToInt(rawParts[0]), s_strings.ToInt(rawParts[1])}

		rawParts = strings.Split(pair[1], "-")
		if len(rawParts) != 2 {
			panic(fmt.Errorf("error parsing second span '%s': must have two numbers split with -", pair[1]))
		}
		s2 := &span.Span{s_strings.ToInt(rawParts[0]), s_strings.ToInt(rawParts[1])}

		if s1.Overlaps(*s2) {
			sum += 1
		}
	}
	return sum
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
