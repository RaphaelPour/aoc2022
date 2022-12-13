package main

import (
	"fmt"
	"encoding/json"
	"sort"

	"github.com/RaphaelPour/stellar/input"
)

const (
	CONTINUE = iota
	GOOD
	BAD
)

type PacketPair struct {
	left, right any
}

func Valid(left, right any) int {
	leftVal, leftValOk := left.(float64)
	rightVal, rightValOk := right.(float64)
	leftList, leftListOk := left.([]any)
	rightList, rightListOk := right.([]any)

	/* 1. both are numbers */
	if leftValOk && rightValOk {
		if leftVal < rightVal {
			return GOOD
		} else if leftVal > rightVal {
			return BAD
		}
		return CONTINUE
	}

	/* 2. both are lists */
	if leftListOk && rightListOk {
		for i := 0; i < len(leftList) && i < len(rightList); i++ {
			if result := Valid(leftList[i], rightList[i]); result != CONTINUE {
				return result
			}
		}

		if len(leftList) < len(rightList) {
			return GOOD
		} else if len(leftList) > len(rightList) {
			return BAD
		}

		return CONTINUE
	}

	/* 3. one is a number and the other a list */
	if leftValOk && rightListOk {
		return Valid([]any{leftVal}, rightList)
	} else if leftListOk && rightValOk {
		return Valid(leftList, []any{rightVal})
	}

	/* else: something went wrong */
	panic(fmt.Sprintf("error asserting '%s' or '%s': expected number or list", left, right))
}

func part1(data []string) int {
	packets := make([]PacketPair, 0)
	sum := 0
	for i, line := range data {
		if i%3 == 0 {
			var left []any
			if err := json.Unmarshal([]byte(line), &left); err != nil {
				panic(fmt.Sprintf("error parsing %s: %s", line, err))
			}
			packets = append(packets, PacketPair{left: left})
		} else if i%3 == 1 {
			var right []any
			if err := json.Unmarshal([]byte(line), &right); err != nil {
				panic(fmt.Sprintf("error parsing %s: %s", line, err))
			}
			packets[len(packets)-1].right = right
		}
	}

	for i, packet := range packets {
		if Valid(packet.left, packet.right) == GOOD{
			sum += i + 1
		}
	}

	return sum
}

func part2(data []string) int {
	packets := make([]any, 0)
	data = append(data,"[[2]]", "[[6]]")
	for _, line := range data {
		if line == ""{
			continue
		}

		var packet []any
		if err := json.Unmarshal([]byte(line), &packet); err != nil {
			panic(fmt.Sprintf("error parsing %s: %s", line, err))
		}
		packets = append(packets, packet)
	}

	divider1 := fmt.Sprintf("%#v",packets[len(packets)-1])
	divider2 := fmt.Sprintf("%#v",packets[len(packets)-2])

	sort.Slice(packets,func(i,j int)bool{
		return Valid(packets[i], packets[j]) == GOOD
	})

	sum := 1
	for i, packet := range packets{
		s := fmt.Sprintf("%#v", packet)
		if s == divider1 || s == divider2{
			sum *= (i+1)
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

	fmt.Println("too low: 16263")
	fmt.Println("too high: 30504")
}
