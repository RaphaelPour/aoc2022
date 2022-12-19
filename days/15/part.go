package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type Range struct {
	start, end int
}

type Rhombus struct {
	conture map[int]Range
}

func (r Rhombus) Contains(other Point) (int, bool) {
	rr, ok := r.conture[other.y]
	if !ok {
		return -1, false
	}
	return rr.end - other.x + 1, other.x >= rr.start && other.x <= rr.end
}

func Dist(a, b Point) int {
	return math.Abs(a.x-b.x) + math.Abs(a.y-b.y)
}

func part1(data []string) int {
	re := regexp.MustCompile(`x=([-\d]+), y=([-\d]+).*x=([-\d]+), y=([-\d]+)`)

	m := make(map[int]struct{})
	for _, line := range data {
		match := re.FindStringSubmatch(line)

		sensor := Point{
			s_strings.ToInt(match[1]),
			s_strings.ToInt(match[2]),
		}
		beacon := Point{
			s_strings.ToInt(match[3]),
			s_strings.ToInt(match[4]),
		}

		dist := Dist(sensor, beacon)
		baseline := Point{sensor.x, 2000000}

		for x := -dist; x <= dist; x++ {
			newP := sensor.Add(Point{x, 0})
			newP.y = baseline.y
			if Dist(newP, sensor) <= dist {
				m[newP.x] = struct{}{}
			}
		}
	}

	return len(m)
}

func part2(data []string, count int) int {
	re := regexp.MustCompile(`x=([-\d]+), y=([-\d]+).*x=([-\d]+), y=([-\d]+)`)

	rhombusList := make([]Rhombus, len(data))
	for i, line := range data {
		match := re.FindStringSubmatch(line)

		sensor := Point{
			s_strings.ToInt(match[1]),
			s_strings.ToInt(match[2]),
		}
		beacon := Point{
			s_strings.ToInt(match[3]),
			s_strings.ToInt(match[4]),
		}

		dist := Dist(sensor, beacon)

		r := Rhombus{}
		r.conture = make(map[int]Range)

		for i := 0; i <= dist; i++ {
			r.conture[sensor.y+i] = Range{
				start: sensor.x - dist + i,
				end:   sensor.x + dist - i,
			}
			r.conture[sensor.y-i] = Range{
				start: sensor.x - dist + i,
				end:   sensor.x + dist - i,
			}
		}

		rhombusList[i] = r
	}

	fmt.Println("scanning")

	for y := 0; y <= count; y++ {
		if y%100000 == 0 {
			fmt.Print(".")
		}
		for x := 0; x <= count; {
			skipped := false
			for _, r := range rhombusList {
				if skip, ok := r.Contains(Point{x, y}); ok {
					// fmt.Println("skip", skip, "to", x+skip)
					x += skip
					skipped = true
					break
				}
			}
			if !skipped {
				fmt.Println(x, y)
				return x*4000000 + y
			}
		}
	}

	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data, 4000000))
}
