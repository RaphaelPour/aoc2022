package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Operation func(a, b int) int

var (
	operations = map[string]Operation{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	cache = map[string]int{}
)

type Monkey struct {
	name          string
	terminal      bool
	number        int
	first, second string
	operation     Operation
}

func NewMonkey(line string) Monkey {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("expected %s to have two parts, got %d", line, len(parts)))
	}

	m := Monkey{}
	m.name = parts[0]

	job := strings.Split(parts[1], " ")
	if len(job) == 1 {
		m.terminal = true
		m.number = s_strings.ToInt(job[0])
	} else if len(job) == 3 {
		m.terminal = false
		m.first = job[0]
		var ok bool
		m.operation, ok = operations[job[1]]
		if !ok {
			panic(fmt.Sprintf("unknown operation %s", job[1]))
		}
		m.second = job[2]
	} else {
		panic(fmt.Sprintf("expected %s to have one or three parts, got %d", parts[1], len(parts)))
	}

	return m
}

func Resolve(start Monkey, monkeys map[string]Monkey) int {
	if start.terminal {
		return start.number
	}

	return start.operation(Resolve(monkeys[start.first], monkeys), Resolve(monkeys[start.second], monkeys))
}

func ResolveWithHuman(start Monkey, num int, monkeys map[string]Monkey, humanMap map[string]struct{}) int {
	if start.name == "humn" {
		return num
	}
	if start.terminal {
		return start.number
	}

	if val, ok := cache[start.name]; ok {
		return val
	}

	var a, b int
	if val, hit := cache[start.first]; hit {
		a = val
	} else {
		a = ResolveWithHuman(monkeys[start.first], num, monkeys, humanMap)

		// cache value if its not on the human path
		if _, found := humanMap[start.first]; !found {
			cache[start.first] = a
		}
	}

	if val, hit := cache[start.second]; hit {
		b = val
	} else {
		b = ResolveWithHuman(monkeys[start.second], num, monkeys, humanMap)

		// cache value if its not on the human path
		if _, found := humanMap[start.second]; !found {
			cache[start.second] = b
		}
	}

	return start.operation(a, b)
}

func HumanPath(start Monkey, monkeys map[string]Monkey) ([]string, bool) {
	if start.name == "humn" {
		return nil, true
	} else if start.terminal {
		return nil, false
	}

	left, ok := HumanPath(monkeys[start.first], monkeys)
	if ok {
		return append([]string{start.name}, left...), true
	}

	right, ok := HumanPath(monkeys[start.second], monkeys)
	if ok {
		return append([]string{start.name}, right...), true
	}

	return nil, false
}

func part1(root Monkey, monkeys map[string]Monkey) int {
	return Resolve(root, monkeys)
}

func part2(root Monkey, monkeys map[string]Monkey) int {
	path, found := HumanPath(root, monkeys)
	if !found {
		panic("path to humn not found")
	}
	pathMap := make(map[string]struct{})
	for _, node := range path {
		pathMap[node] = struct{}{}
	}

	for i := 0; ; i++ {
		if (i+1)%1000000 == 0 {
			fmt.Println(i)
		}
		a := ResolveWithHuman(monkeys[root.first], i, monkeys, pathMap)
		b := ResolveWithHuman(monkeys[root.second], i, monkeys, pathMap)
		if a == b {
			fmt.Println("EQUALIY!", a, b)
			return i
		}
	}
	return -1
}

func main() {
	monkeys := make(map[string]Monkey)
	var root Monkey
	for _, line := range input.LoadString("input1") {
		m := NewMonkey(line)
		monkeys[m.name] = m

		if m.name == "root" {
			root = m
		}
	}

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(root, monkeys))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(root, monkeys))

	fmt.Println("checked til 1871999999")
}
