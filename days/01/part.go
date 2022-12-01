package main

import (
	"fmt"
	"sort"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/strings"
)

func part1(data []string) int {
	maxCalories := 0
	currentCalories := 0
	for i, row := range data {
		if row == "" || i == len(data)-1 {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}
		currentCalories += strings.ToInt(row)
	}
	return maxCalories
}

func part2(data []string) int {
	calories := make([]int, 0)
	currentCalories := 0
	for i, row := range data {
		if row == "" || i == len(data)-1 {
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		currentCalories += strings.ToInt(row)
	}

	sort.Ints(calories)

	l := len(calories)
	return calories[l-1] + calories[l-2] + calories[l-3]
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(data))
}
