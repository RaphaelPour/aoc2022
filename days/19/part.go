package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

var (
	robotMap = map[string]Material{
		"ore":      Material{ore: 1},
		"clay":     Material{clay: 1},
		"obsidian": Material{obsidian: 1},
		"geode":    Material{geode: 1},
	}
)

type CacheKey struct {
	cost, robots Material
	minutesLeft  int
}

type Material struct {
	ore, clay, obsidian, geode int
}

func (m Material) String() string {
	return fmt.Sprintf(
		"ore=%d clay=%d obsidian=%d geode=%d",
		m.ore, m.clay, m.obsidian, m.geode,
	)
}

func (stock Material) IsAffordable(cost Material) bool {
	return stock.ore >= cost.ore &&
		stock.clay >= cost.clay &&
		stock.obsidian >= cost.obsidian
}

func (m Material) Add(other Material) Material {
	m.ore += other.ore
	m.clay += other.clay
	m.obsidian += other.obsidian
	m.geode += other.geode
	return m
}

func (m Material) Sub(other Material) Material {
	m.ore -= other.ore
	m.clay -= other.clay
	m.obsidian -= other.obsidian
	m.geode -= other.geode
	return m
}

type Blueprint struct {
	materials map[string]Material
	cache     map[CacheKey]int
}

func (b Blueprint) Next(stock, robots Material, minutesLeft int) int {
	// exit recursion if time has run out
	if minutesLeft <= 0 {
		return stock.geode
	}

	// collect
	stock = stock.Add(robots)

	if geodes, ok := b.cache[CacheKey{stock, robots, minutesLeft}]; ok {
		return geodes
	}

	// divide and conquer on buying robots
	maxGeodes := b.Next(stock, robots, minutesLeft-1)

	for material, cost := range b.materials {
		if !stock.IsAffordable(cost) {
			continue
		}
		if geodes := b.Next(stock.Sub(cost), robots.Add(robotMap[material]), minutesLeft-1); geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if maxGeodes > 12 {
		panic(maxGeodes)
	}

	if maxGeodes == 0 {
		b.cache[CacheKey{stock, robots, minutesLeft - 1}] = 0
		if len(b.cache)%1000000 == 0 {
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
			materials: map[string]Material{
				"ore":  Material{ore: s_strings.ToInt(match[1][1])},
				"clay": Material{ore: s_strings.ToInt(match[2][1])},
				"obsidian": Material{
					ore:  s_strings.ToInt(match[3][1]),
					clay: s_strings.ToInt(match[4][1]),
				},
				"geode": Material{
					ore:      s_strings.ToInt(match[5][1]),
					obsidian: s_strings.ToInt(match[6][1]),
				},
			},
		}
		b.cache = make(map[CacheKey]int)
		sum += (i + 1) * b.Next(Material{}, Material{ore: 1}, 24)
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
