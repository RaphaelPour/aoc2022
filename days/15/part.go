package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Map struct {
	beacons map[math.Point[int]]struct{}
	sensors map[math.Point[int]]struct{}
}

func part1(data []string) int {
	re := regexp.MustCompile(`x=([-\d]+), y=([-\d]+).*x=([-\d]+), y=([-\d]+)`)

	m := Map{}
	m.beacons = make(map[math.Point[int]]struct{})
	m.sensors = make(map[math.Point[int]]struct{})
	for _, line := range data {
		match := re.FindStringSubmatch(line)
		sx := s_strings.ToInt(match[1])
		sy := s_strings.ToInt(match[2])
		bx := s_strings.ToInt(match[3])
		by := s_strings.ToInt(match[4])

		m.beacons[math.Point[int]{bx, by}] = struct{}{}
		m.sensors[math.Point[int]{sx, sy}] = struct{}{}
	}

	fmt.Println(m)
	return 0
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
