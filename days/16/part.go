package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Valve struct {
	name      string
	flowRate  int
	neighbors []string
}

type CacheKey struct{
	start *Valve
	open string
}
type CacheValue struct{
	pressure int
	path string
}

type Problem struct {
	valvesWithFlow int
	valves         map[string]Valve
	cache          map[CacheKey]CacheValue
}

func (p *Problem) Find(start Valve, minutesLeft int, releasedPressure int, open string) (int, string) {
	// return if all valves with flow have been visited
	if len(open) >= p.valvesWithFlow*3 {
		return releasedPressure, fmt.Sprintf("%s (%dmin), all visited", start.name, 30-minutesLeft)
	}

	// return if time ran out
	if minutesLeft <= 0 {
		return releasedPressure, fmt.Sprintf("%s (%dmin)", start.name, 30-minutesLeft)
	}

	// stay one minute for releasing pressure if the valve is still closed and has flow rate 
	if start.flowRate > 0 && !strings.Contains(open, start.name) {
		minutesLeft--
		releasedPressure += minutesLeft * start.flowRate
		open += fmt.Sprintf(" %s", start.name)
	}

	max := 0
	currentPath := ""
	for _, neighbor := range start.neighbors {
		if val,ok := p.cache[CacheKey{&start, open}];ok {
			if max < val.pressure{
				max = val.pressure
				currentPath = val.path
				continue
			}
		}
		pressure, path := p.Find(p.valves[neighbor], minutesLeft-1, releasedPressure, open)
		if pressure > max {
			currentPath = path
			max = pressure
		}
		p.cache[CacheKey{&start, open}] = CacheValue{pressure, path}
	}

	return max, fmt.Sprintf("%s (%dmin, %d), %s", start.name, 30-minutesLeft, releasedPressure, currentPath)
}

func part1(data []string) int {
	re := regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); tunnel(s)? lead(s)? to valve(s)? (.*)$`)
	p := new(Problem)
	p.valves = make(map[string]Valve)
	p.cache = make(map[CacheKey]CacheValue)
	for _, line := range data {
		match := re.FindStringSubmatch(line)
		p.valves[match[1]] = Valve{
			name:      match[1],
			flowRate:  s_strings.ToInt(match[2]),
			neighbors: strings.Split(strings.ReplaceAll(match[6], " ", ""), ","),
		}

		if p.valves[match[1]].flowRate > 0 {
			p.valvesWithFlow++
		}
	}

	pressure, path := p.Find(p.valves["AA"], 30, 0, "")
	fmt.Println(path)
	return pressure
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
