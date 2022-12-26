package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

var (
	snafuMap = map[string]int{
		"=": -2,
		"-": -1,
		"0": 0,
		"1": 1,
		"2": 2,
	}

	decMap = map[int]string{
		-2: "=",
		-1: "-",
		0:  "0",
		1:  "1",
		2:  "2",
	}
)

func snafu2Dec(n string) int {
	result := 0
	for i := len(n) - 1; i >= 0; i-- {
		result += math.Pow(5, len(n)-i-1) * snafuMap[string(n[i])]
	}
	return result
}

func dec2Snafu(n int) string {
	result := ""
	translate := []string{"0", "1", "2", "=", "-"}
	for n != 0 {
		m := n % 5
		n /= 5
		if m > 2 {
			n++
		}
		result = translate[m] + result
	}
	return result
}

func part1(data []string) string {
	sum := 0
	for _, line := range data {
		fmt.Println(line)
		sum += snafu2Dec(line)
	}
	fmt.Println(sum)
	return dec2Snafu(sum)
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
