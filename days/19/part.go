package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	robotMap = map[string]Cost{
		"ore": Cost{ore:1},
		"clay":Cost{clay:1},
		"obsidian":Cost{obsidian:1},
		"geode":Cost{geode:1},
	}
)

type CacheKey struct {
	cost, robots Cost
	minutesLeft int
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
	materials map[string]Cost
	cache map[CacheKey]int
}

func (b Blueprint) Next(stock, robots Cost, minutesLeft int) int {
	// exit recursion if time has run out
	if minutesLeft <= 0 {
		return stock.geode
	}

	// collect
	stock.Add(robots)

	if geodes, ok := b.cache[CacheKey{stock, robots, minutesLeft}];ok{
		return geodes
	}

	// divide and conquer on buying robots
	maxGeodes := b.Next(stock, robots, minutesLeft-1)

	for material, cost := range b.materials {
		if !stock.IsAffordable(cost) {
			continue
		}
		if geodes := b.Next(stock.SubNew(cost), robots.AddNew(robotMap[material]),minutesLeft-1); geodes > maxGeodes{
			maxGeodes = geodes
		}
	}

	if maxGeodes > 12 {
		panic(maxGeodes)
	}

	if maxGeodes > 10 {
		b.cache[CacheKey{stock, robots, minutesLeft-1}] = maxGeodes
		if len(b.cache) % 10000 == 0 {
			fmt.Println(len(b.cache))
		}
	}

	return maxGeodes
}

func part1(data []string) int {
	re := regexp.MustCompile(`(\d+)`)
	sum := 0
	for i, line := range data {
		match := re.FindAllStringSubmatch(line, -1)

		b := Blueprint{
			materials: map[string]Cost{
				"ore": Cost{ore: s_strings.ToInt(match[1][1])},
				"clay": Cost{ore: s_strings.ToInt(match[2][1])},
				"obsidian": Cost{
					ore:  s_strings.ToInt(match[3][1]),
					clay: s_strings.ToInt(match[4][1]),
				},
				"geode": Cost{
					ore:      s_strings.ToInt(match[5][1]),
					obsidian: s_strings.ToInt(match[6][1]),
				},
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
