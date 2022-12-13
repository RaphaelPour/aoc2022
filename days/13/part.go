package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

const (
	CONTINUE = iota
	BAD
	GOOD
)

type PacketPair struct {
	left, right string
}

func Valid(left, right string) int {
	fmt.Println("\nCompare", left, "vs", right)
	/* 1. check if both are integers */
	leftVal, err1 := strconv.Atoi(left)
	rightVal, err2 := strconv.Atoi(right)
 
	if err1 == nil && err2 == nil {
		fmt.Println("both numbers")
		if leftVal < rightVal {
				fmt.Println("good")
				return GOOD
		} else if rightVal < leftVal {
				fmt.Println("bad")
			return BAD
		}
		fmt.Println("continue")
		return CONTINUE
	}

	/* 2. both are lists */
	if err1 != nil && err2 != nil {
		fmt.Println("both lists")
		leftList := strings.Split(left[1:len(left)-1], ",")
		rightList := strings.Split(right[1:len(right)-1], ",")

		fmt.Printf("%s -> %#v\n",left, leftList)
		fmt.Printf("%s -> %#v\n",right,  rightList)
		// return false

		/* compare pair-wise */
		for i := range leftList {
			if result := Valid(leftList[i], rightList[i]); result != CONTINUE {
				fmt.Println("compare",result)
				return result
			}
		}
		
		if len(rightList) > len(leftList) {
			fmt.Println("good")
			return GOOD
		}
		
		/* right list shouldn't have less than the left one */
		if len(rightList) < len(leftList) {
				fmt.Println("bad")
			return BAD
		}

		fmt.Println("continue")
		return CONTINUE
	}

	fmt.Println("mixed")
	return BAD
	/* 3. one of both is an integer */
	if err1 != nil {
		left = fmt.Sprintf("[%s]", left)
	}
	if err2 != nil {
		right = fmt.Sprintf("[%s]", right)
	}
	return Valid(left, right)
}

func part1(data []string) int {
	packets := make([]PacketPair, 0)
	sum := 0
	for i, line := range data {
		if i%3 == 0 {
			packets = append(packets, PacketPair{left: line})
		} else if i%3 == 1 {
			packets[len(packets)-1].right = line
		} else {
			p := packets[len(packets)-1]
			result := Valid(p.left, p.right)
			if result == GOOD {
				sum += len(packets) - 1 + 1
			}
		}

		if i == 5 {
			return sum
		}
	}

	return sum
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))
}
