package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

type Window struct {
	data   []string
	length int
}

func (w *Window) Add(item string) {
	fmt.Println(w.data)
	if len(w.data) >= w.length {
		w.data = w.data[1:len(w.data)]
	}

	w.data = append(w.data, item)
}

func (w *Window) IsUnique() bool {
	if len(w.data) < w.length {
		return false
	}

	for i := 0; i < len(w.data)-1; i++ {
		for j := i + 1; j < len(w.data); j++ {
			if i == j {
				continue
			}
			if w.data[i] == w.data[j] {
				return false
			}
		}
	}
	return true
}

func part1(data string) int {
	w := Window{
		make([]string, 0),
		4,
	}
	for i := 0; i < len(data); i++ {
		w.Add(string(data[i]))
		if w.IsUnique() {
			return i + 1
		}
	}
	return 0
}

func part2(data string) int {
	w := Window{
		make([]string, 0),
		14,
	}
	for i := 0; i < len(data); i++ {
		w.Add(string(data[i]))
		if w.IsUnique() {
			return i + 1
		}
	}
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data[0]))

	fmt.Println("bad: 1537")

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data[0]))
}
