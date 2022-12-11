package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	m_math "github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	modus = 1
)

type Operation struct {
	raw      string
	operator func(int, int) int
	operand  string
}

func NewOperation(input string) Operation {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 2 {
		panic(fmt.Sprintf("error parsing operation '%s': expected two fields, got %d", input, len(parts)))
	}

	o := Operation{}
	o.raw = input
	o.operand = parts[1]

	switch parts[0] {
	case "+":
		o.operator = func(a, b int) int { return a + b }
	case "-":
		o.operator = func(a, b int) int { return a - b }
	case "*":
		o.operator = func(a, b int) int { return a * b }
	case "/":
		o.operator = func(a, b int) int { return a / b }
	default:
		panic(fmt.Sprintf("unknown operation %s", parts[0]))
	}

	return o
}

func (o Operation) Compute(old int) int {
	var other int

	if o.operand == "old" {
		other = old
	} else {
		other = s_strings.ToInt(o.operand)
	}

	return o.operator(old, other)
}

type Monkey struct {
	items         []int
	operation     Operation
	divisibleBy   int
	ifTrueTarget  int
	ifFalseTarget int

	autoWorryRelief bool
	inspected       int
}

func NewMonkey(autoWorryRelief bool) *Monkey {
	m := new(Monkey)
	m.items = make([]int, 0)
	m.autoWorryRelief = autoWorryRelief
	return m
}

func (m Monkey) Test(worryLevel int) int {
	m.inspected += 1
	if worryLevel%m.divisibleBy == 0 {
		return m.ifTrueTarget
	}

	return m.ifFalseTarget
}

type Monkeys struct {
	bullies []*Monkey
	modues  int
}

func NewMonkeys() *Monkeys {
	m := new(Monkeys)
	m.bullies = make([]*Monkey, 0)
	m.modues = 1
	return m
}

func (m *Monkeys) Add(monkey *Monkey) {
	m.bullies = append(m.bullies, monkey)
	m.modues *= monkey.divisibleBy
}

func (m Monkeys) Round() {
	for _, monkey := range m {
		// fmt.Printf("Monkey %d:\n", i)
		for _, item := range monkey.items {
			// fmt.Printf("\tMonkey inspects an item withj a worry level of %d.\n", item)

			// monkey inspects item -> operation gets applied
			// fmt.Printf("\t\tWorry level = worry level %s\n", monkey.operation.raw)
			item = monkey.operation.Compute(item) % modus

			// monkey gets bored -> divide item by 3
			if monkey.autoWorryRelief {
				item /= 3
			}
			// fmt.Printf("\t\tMonkey gets bored: Worry level is divided by 3 to %d.\n", item)

			// test item -> throw it to the resulting monkey
			target := monkey.Test(item)
			// fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", item, target)
			m[target].items = append(m[target].items, item)
			monkey.inspected += 1
		}

		// monkey has thrown all items away
		monkey.items = monkey.items[:0]
	}
}

func part1(data []string) int {
	re := regexp.MustCompile(`^Monkey.*Starting items:([\s\d,]+).*new = old (.*).*Test: divisible by (\d+) .*monkey (\d+).*monkey (\d+)`)

	monkeys := NewMonkeys()
	for _, rawLine := range data {
		line := strings.ReplaceAll(rawLine, "\n", " ")
		match := re.FindStringSubmatch(line)

		monkey := NewMonkey(true)
		rawItems := strings.Split(match[1], ",")
		for _, item := range rawItems {
			monkey.items = append(monkey.items, s_strings.ToInt(strings.TrimSpace(item)))
		}

		monkey.operation = NewOperation(match[2])
		monkey.divisibleBy = s_strings.ToInt(match[3])
		monkey.ifTrueTarget = s_strings.ToInt(match[4])
		monkey.ifFalseTarget = s_strings.ToInt(match[5])

		monkeys.Add(monkey)
	}

	for round := 1; round <= 20; round++ {
		// fmt.Println("round:", round)
		monkeys.Round()
	}

	inspections := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspections[i] = monkey.inspected
	}

	return m_math.Product(m_math.MaxN(inspections, 2))
}

/*
func part2(data []string) int {
	re := regexp.MustCompile(`^Monkey.*Starting items:([\s\d,]+).*new = old (.*).*Test: divisible by (\d+) .*monkey (\d+).*monkey (\d+)`)

	monkeys := make(Monkeys, 0)
	for _, rawLine := range data {
		line := strings.ReplaceAll(rawLine, "\n", " ")
		match := re.FindStringSubmatch(line)
		// fmt.Printf("%#v\n\n", match)

		monkey := NewMonkey(false)
		rawItems := strings.Split(match[1], ",")
		for _, item := range rawItems {
			monkey.items = append(monkey.items, s_strings.ToInt(strings.TrimSpace(item)))
		}

		monkey.operation = NewOperation(match[2])
		monkey.divisibleBy = s_strings.ToInt(match[3])
		monkey.ifTrueTarget = s_strings.ToInt(match[4])
		monkey.ifFalseTarget = s_strings.ToInt(match[5])

		monkeys = append(monkeys, monkey)
	}

	for round := 1; round <= 10000; round++ {
		// fmt.Println("round:", round)
		monkeys.Round()
	}

	inspections := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspections[i] = monkey.inspected
	}

	return m_math.Product(m_math.MaxN(inspections, 2))
}
*/

func LoadStringWithDelimiter(filename, delimiter string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	return strings.Split(string(data), delimiter)
}

func main() {
	data := LoadStringWithDelimiter("input", "\n\n")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	/*
			 * fmt.Println("== [ PART 2 ] ==")
		     * fmt.Println(part2(data))
	*/
}
