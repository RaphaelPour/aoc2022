package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type Cost struct {
	ore, clay, obsidian, geode int
}

func (c Cost) String() string {
	return fmt.Sprintf(
		"ore=%d clay=%d obsidian=%d geode=%d",
		c.ore, c.clay, c.obsidian, c.geode,
	)
}

func (stock Cost) IsAffordable(cost Cost) bool {
	return stock.ore >= cost.ore &&
	       stock.clay >= cost.clay &&
				 stock.obsidian >= cost.obsidian &&
				 stock.geode >= cost.geode
}

func (c *Cost) Add(other Cost) {
	c.ore += other.ore
	c.clay += other.clay
	c.obsidian += other.obsidian
	c.geode += other.geode
}

func (c *Cost) Sub(other Cost) {
	c.ore -= other.ore
	c.clay -= other.clay
	c.obsidian -= other.obsidian
	c.geode -= other.geode
}

func (c *Cost) Buy(other Cost) {
	c.Sub(other)
}

type Blueprint struct {
	ore, clay, obsidian, geode Cost
}

func (b Blueprint) Do(stock,robots Cost, minutesLeft int) int {
	for minutes := minutesLeft; minutes > 0; minutes-- {
		fmt.Printf("== Minute %d ==\n", 24- minutes)
		queue := make([]Cost, 0)
		// Try build stuff starting with geode
		if stock.IsAffordable(b.geode) {
			stock.Buy(b.geode)
			queue = append(queue, Cost{geode: 1})
			fmt.Printf("Spend %s to start building a geode robot.\n", b.geode)
		}

		if stock.IsAffordable(b.obsidian) {
			stock.Buy(b.obsidian)
			queue = append(queue, Cost{obsidian: 1})
			fmt.Printf("Spend %s to start building a obsidian robot.\n", b.obsidian)
		}

		if stock.IsAffordable(b.clay) {
			stock.Buy(b.clay)
			queue = append(queue, Cost{clay: 1})
			fmt.Printf("Spend %s to start building a clay robot.\n", b.clay)
		}

		if stock.IsAffordable(b.ore) {
			stock.Buy(b.ore)
			queue = append(queue, Cost{ore: 1})
			fmt.Printf("Spend %s to start building a ore robot.\n", b.ore)
		}

		// Collect
		if robots.ore > 0 {
			stock.ore += robots.ore
			fmt.Printf(
				"%d ore-collecting robot collects %d ore; you now have %d ore.\n",
				robots.ore, robots.ore, stock.ore,
			)
		}

		if robots.clay > 0 {
			stock.clay += robots.clay
			fmt.Printf(
				"%d clay-collecting robot collects %d clay; you now have %d clay.\n",
				robots.clay, robots.clay, stock.clay,
			)
		}

		if robots.obsidian > 0 {
			stock.obsidian += robots.obsidian
			fmt.Printf(
				"%d obsidian-collecting robot collects %d obsidian; you now have %d obsidian.\n",
				robots.obsidian, robots.obsidian, stock.obsidian,
			)
		}

		if robots.geode > 0 {
			stock.geode += robots.geode
			fmt.Printf(
				"%d geode-cracking robot collects %d geode; you now have %d geode.\n",
				robots.geode, robots.geode, stock.geode,
			)
		}

		for _, robot := range queue {
			robots.Add(robot)
			fmt.Printf("New robot %s arrived\n", robot)
		}
	}
	fmt.Println(stock)
	return stock.geode
}

func part1(data []string) int {
	re := regexp.MustCompile(`(\d+)`)
	for _, line := range data {
		match := re.FindAllStringSubmatch(line, -1)

		b := Blueprint{
			ore:  Cost{ore: s_strings.ToInt(match[1][1])},
			clay: Cost{ore: s_strings.ToInt(match[2][1])},
			obsidian: Cost{
				ore:  s_strings.ToInt(match[3][1]),
				clay: s_strings.ToInt(match[4][1]),
			},
			geode: Cost{
				ore:      s_strings.ToInt(match[5][1]),
				obsidian: s_strings.ToInt(match[6][1]),
			},
		}

		fmt.Println(b.Do(Cost{}, Cost{ore:1},24))
		break
	}

	return 0
}

func part2(data []string) int {
	return 0
}

func main() {
	data := input.LoadString("input1")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
