package main

import (
	"fmt"
	"github.com/RaphaelPour/stellar/input"
	"strings"
)

func priority(c rune) int {
	if int(c) >= 'a' && int(c) <= 'z' {
		return int(c-'a') + 1
	}
	return (int(c) - 'A') + 27
}

type Rucksack struct {
	total       string
	left, right string
}

func NewRucksack(total string) *Rucksack {
	r := new(Rucksack)
	r.total = total
	r.left = total[0 : len(total)/2]
	r.right = total[len(total)/2:]
	return r
}

func (r Rucksack) Duplicate() (int, bool) {
	for _, c := range r.left {
		if strings.Contains(r.right, string(c)) {
			return priority(c), true
		}
	}
	return -1, false
}

func (r Rucksack) DuplicatesWithOther(other Rucksack) []rune {
	result := make([]rune, 0)
	for _, c := range r.total {
		if strings.Contains(other.total, string(c)) {
			result = append(result, c)
		}
	}
	return result
}

type Group []*Rucksack

func (g Group) Duplicate() (int, bool) {
	for _, d := range g[0].DuplicatesWithOther(*g[1]) {
		if strings.Contains(g[2].total, string(d)) {
			return priority(d), true
		}
	}
	return -1, false
}

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		r := NewRucksack(line)
		if priority, ok := r.Duplicate(); ok {
			sum += priority
		}
	}
	return sum
}

func part2(data []string) int {
	sum := 0
	groups := make([]Group, 0)
	for i, line := range data {
		r := NewRucksack(line)
		if i%3 == 0 {
			groups = append(groups, Group{r})
		} else {
			groups[len(groups)-1] = append(groups[len(groups)-1], r)
		}
	}

	for _, g := range groups {
		if p, ok := g.Duplicate(); ok {
			sum += p
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
