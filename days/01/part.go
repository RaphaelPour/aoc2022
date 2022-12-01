package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	"github.com/RaphaelPour/stellar/strings"
)

func maxCalories(data []string, n int) int {
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

	return math.Sum(math.MaxN(calories, n))
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(maxCalories(data, 1))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(maxCalories(data, 3))
}
