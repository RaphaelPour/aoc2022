package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	valvesWithFlow int
)

type Valve struct {
	name      string
	flowRate  int
	neighbors []string
}

func Find(start Valve, minutesLeft int, releasedPressure int, open string, valves map[string]Valve) (int, string) {
	if len(open) == valvesWithFlow*3 {
		return releasedPressure, fmt.Sprintf("%s (%dmin), all visited", start.name, 30-minutesLeft)
	}

	if minutesLeft <= 0 {
		return releasedPressure, fmt.Sprintf("%s (%dmin)", start.name, 30-minutesLeft)
	}

	if start.flowRate > 0 && !strings.Contains(open, start.name) {
		minutesLeft--
		releasedPressure += minutesLeft * start.flowRate
		open += fmt.Sprintf(" %s", start.name)
	}

	max := 0
	currentPath := ""
	for _, neighbor := range start.neighbors {
		pressure, path := Find(valves[neighbor], minutesLeft-1, releasedPressure, open, valves)
		if pressure > max {
			currentPath = path
			max = pressure
		}
	}

	return max, fmt.Sprintf("%s (%dmin, %d), %s", start.name, 30-minutesLeft, releasedPressure, currentPath)
}

func part1(data []string) int {
	re := regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); tunnel(s)? lead(s)? to valve(s)? (.*)$`)
	m := make(map[string]Valve)
	for _, line := range data {
		match := re.FindStringSubmatch(line)
		m[match[1]] = Valve{
			name:      match[1],
			flowRate:  s_strings.ToInt(match[2]),
			neighbors: strings.Split(strings.ReplaceAll(match[6], " ", ""), ","),
		}

		if m[match[1]].flowRate > 0 {
			valvesWithFlow++
		}
	}

	pressure, path := Find(m["AA"], 30, 0, "", m)
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
