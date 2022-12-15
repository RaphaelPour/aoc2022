package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Map struct {
	blocked  map[math.Point[int]]struct{}
	count    int
	min, max math.Point[int]
}

func NewMap() *Map {
	m := new(Map)
	m.blocked = make(map[math.Point[int]]struct{})
	m.min = math.Point[int]{10000, 10000}
	return m
}

func (m *Map) Add(p math.Point[int]) {
	if _, ok := m.blocked[p]; ok {
		return
	}

	m.blocked[p] = struct{}{}
	m.count++
}

func Dist(a, b math.Point[int]) int {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

func part1(data []string) int {
	re := regexp.MustCompile(`x=([-\d]+), y=([-\d]+).*x=([-\d]+), y=([-\d]+)`)

	m := make(map[int]struct{})
	for _, line := range data {
		match := re.FindStringSubmatch(line)

		sensor := math.Point[int]{
			s_strings.ToInt(match[1]),
			s_strings.ToInt(match[2]),
		}
		beacon := math.Point[int]{
			s_strings.ToInt(match[3]),
			s_strings.ToInt(match[4]),
		}

		dist := Dist(sensor, beacon)
		baseline := math.Point[int]{sensor.X, 2000000}
		// baseline := math.Point[int]{sensor.X, 10}

		for x := -dist; x <= dist; x++ {
			newP := sensor.Add(math.Point[int]{x, 0})
			newP.Y = baseline.Y
			if Dist(newP, sensor) <= dist { // && newP != beacon{
				m[newP.X] = struct{}{}
			}
		}
	}

	return len(m)
}

func part2(data []string) int {
	re := regexp.MustCompile(`x=([-\d]+), y=([-\d]+).*x=([-\d]+), y=([-\d]+)`)

	m := make(map[math.Point[int]]struct{})
	for _, line := range data {
		fmt.Println(line)
		match := re.FindStringSubmatch(line)

		sensor := math.Point[int]{
			s_strings.ToInt(match[1]),
			s_strings.ToInt(match[2]),
		}
		beacon := math.Point[int]{
			s_strings.ToInt(match[3]),
			s_strings.ToInt(match[4]),
		}

		dist := Dist(sensor, beacon)+1

		for y := -dist; y <= dist; y++{
			for x := -dist; x <= dist; x++ {
				newP := sensor.Add(math.Point[int]{x, y})

				d := Dist(newP, sensor)
				if _, ok := m[newP]; ok  && d < dist {
					delete(m, newP)
				} else if d == dist { // && newP != beacon{
					m[newP] = struct{}{}
				}
			}
		}
	}

	if len(m) != 1 {
		panic(fmt.Sprintf("expected map to have one element, got %d", len(m)))
	}

	for key := range m {
		return key.X * 4000000 + key.Y
	}

	// unreachable code
	return 0
}

func main() {
	data := input.LoadString("input1")

	// fmt.Println("== [ PART 1 ] ==")
	// fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
