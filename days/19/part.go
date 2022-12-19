package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type CacheKey struct {
	cost, robots Cost
}

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

func (c Cost) AddNew(other Cost) Cost {
	c.ore += other.ore
	c.clay += other.clay
	c.obsidian += other.obsidian
	c.geode += other.geode
	return c
}

func (c Cost) SubNew(other Cost) Cost{
	c.ore -= other.ore
	c.clay -= other.clay
	c.obsidian -= other.obsidian
	c.geode -= other.geode
	return c
}

func (c *Cost) Buy(other Cost) {
	c.Sub(other)
}

type Blueprint struct {
	ore, clay, obsidian, geode Cost
	cache map[CacheKey]int
}

func (b Blueprint) Next(stock, robots Cost, minutesLeft int) int {
	// exit recursion if time has run out
	if minutesLeft <= 0 {
		return stock.geode
	}

	fmt.Println(24 - minutesLeft + 1)

	// collect
	stock.Add(robots)

	// divide and conquer on buying robots
	maxGeodes := 0
	if stock.IsAffordable(b.geode) {
		fmt.Println(stock)
		if geodes := b.Next(stock.SubNew(b.geode), robots.AddNew(Cost{geode:1}),minutesLeft-1); geodes > maxGeodes{
			maxGeodes = geodes
		}
	}

	if stock.IsAffordable(b.obsidian) {
		if geodes := b.Next(stock.SubNew(b.obsidian), robots.AddNew(Cost{obsidian:1}), minutesLeft-1); geodes > maxGeodes{
			maxGeodes = geodes
		}
	}

	if stock.IsAffordable(b.clay) {
		if geodes := b.Next(stock.SubNew(b.clay), robots.AddNew(Cost{clay:1}),minutesLeft-1); geodes > maxGeodes{
			maxGeodes = geodes
		}
	}

	if stock.IsAffordable(b.ore) {
		if geodes := b.Next(stock.SubNew(b.ore), robots.AddNew(Cost{ore:1}),minutesLeft-1); geodes > maxGeodes{
			maxGeodes = geodes
		}
	}

	if geodes := b.Next(stock, robots, minutesLeft-1); geodes > maxGeodes{
		maxGeodes = geodes
	}

	if maxGeodes > 12 {
		panic(fmt.Sprintf("%d geodes are too much", maxGeodes))
	}


	return maxGeodes
}

func part1(data []string) int {
	re := regexp.MustCompile(`(\d+)`)
	sum := 0
	for i, line := range data {
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
		b.cache = make(map[CacheKey]int)

		sum += (i+1) * b.Next(Cost{}, Cost{ore:1},24)
		break
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

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
