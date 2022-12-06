package main

import (
	"regexp"
	"fmt"
	"strings"
	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
	"github.com/RaphaelPour/stellar/stack"
	
)

func part1(data []string) int {
	/* Example regex match
	 * 
     *    [D]    : [][]string{[]string{"    ", "   ", ""}, []string{"[D] ", "[D]", "D"}, []string{"   ", "   ", ""}}
     * [N] [C]    : [][]string{[]string{"[N] ", "[N]", "N"}, []string{"[C] ", "[C]", "C"}, []string{"   ", "   ", ""}}
     * [Z] [M] [P]: [][]string{[]string{"[Z] ", "[Z]", "Z"}, []string{"[M] ", "[M]", "M"}, []string{"[P]", "[P]", "P"}}
	 */
	linePattern := regexp.MustCompile(`(\s\s\s|\[([A-Z])\])\s?`)
	var stacks []stack.Stack[string]
	var skip int
	for i, line := range data {
		if strings.HasPrefix(line," 1") {
			// stack input finished
			// add one to the skip for the empty line after the stack indices
			skip = i+2
			break
		}
	
		m := linePattern.FindAllStringSubmatch(line,-1)
		if len(stacks) == 0 {
			// create all stacks at the first line,
			// the regex should result ALWAYS with the same amount
			// of results but some results may be empty.
			stacks = make([]stack.Stack[string],len(m))
		} else if len(m) == 0 {
			// we reached the 
		} else if len(m) != len(stacks){
			// otherwise the match count MUST equal the count of stacks
			// if not, the regexp didn't match properly
			panic(fmt.Sprintf(
				"error parsing line %s: expected %d matches, got %d",
				line, len(stacks), len(m),
			))
		}

		for i, match := range m {
			item := match[2]
			if item == "" {
				// skip empty item, they needed to be matched in the first place
				// so actual items get pushed into the right stack
				continue
			}

			stacks[i].PushAhead(item)
		}
	}

	// parse moves
	movePattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range data[skip:] {
		m := movePattern.FindStringSubmatch(line)
		if len(m) != 4{
			panic(fmt.Sprintf("error matching %s: expected 4 matches, got %d", line, len(m)))
		}
	
		moves, from, to := s_strings.ToInt(m[1]), s_strings.ToInt(m[2]), s_strings.ToInt(m[3])
		for i := 0; i < moves; i++ {
			// sub 1 since from/to are positions and must be indices
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}


	// Show what's on top of the stack
	for _, stack := range stacks {
		fmt.Print(stack.Pop())
	}
	fmt.Println("")

	return 0
}

func part2(data []string) int {
	/* Example regex match
	 * 
     *    [D]    : [][]string{[]string{"    ", "   ", ""}, []string{"[D] ", "[D]", "D"}, []string{"   ", "   ", ""}}
     * [N] [C]    : [][]string{[]string{"[N] ", "[N]", "N"}, []string{"[C] ", "[C]", "C"}, []string{"   ", "   ", ""}}
     * [Z] [M] [P]: [][]string{[]string{"[Z] ", "[Z]", "Z"}, []string{"[M] ", "[M]", "M"}, []string{"[P]", "[P]", "P"}}
	 */
	linePattern := regexp.MustCompile(`(\s\s\s|\[([A-Z])\])\s?`)
	var stacks []stack.Stack[string]
	var skip int
	for i, line := range data {
		if strings.HasPrefix(line," 1") {
			// stack input finished
			// add one to the skip for the empty line after the stack indices
			skip = i+2
			break
		}
	
		m := linePattern.FindAllStringSubmatch(line,-1)
		if len(stacks) == 0 {
			// create all stacks at the first line,
			// the regex should result ALWAYS with the same amount
			// of results but some results may be empty.
			stacks = make([]stack.Stack[string],len(m))
		} else if len(m) == 0 {
			// we reached the 
		} else if len(m) != len(stacks){
			// otherwise the match count MUST equal the count of stacks
			// if not, the regexp didn't match properly
			panic(fmt.Sprintf(
				"error parsing line %s: expected %d matches, got %d",
				line, len(stacks), len(m),
			))
		}

		for i, match := range m {
			item := match[2]
			if item == "" {
				// skip empty item, they needed to be matched in the first place
				// so actual items get pushed into the right stack
				continue
			}

			stacks[i].PushAhead(item)
		}
	}

	// parse moves
	movePattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range data[skip:] {
		m := movePattern.FindStringSubmatch(line)
		if len(m) != 4{
			panic(fmt.Sprintf("error matching %s: expected 4 matches, got %d", line, len(m)))
		}
	
		moves, from, to := s_strings.ToInt(m[1]), s_strings.ToInt(m[2]), s_strings.ToInt(m[3])
		stacks[to-1].PushN(stacks[from-1].PopN(moves))
	}


	// Show what's on top of the stack
	for _, stack := range stacks {
		fmt.Print(stack.Pop())
	}
	fmt.Println("")

	return 0
}

func main() {
	data := input.LoadString("input")
    
    fmt.Println("== [ PART 1 ] ==")
    fmt.Println(part1(data))

    fmt.Println("== [ PART 2 ] ==")
    fmt.Println(part2(data))
}
