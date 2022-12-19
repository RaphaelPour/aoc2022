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

func (b *Blueprint) Cache(stock, robots Material, geodes, minutesLeft int) {
	/*if minutesLeft > 20 {
		return
	}*/

	if val, ok := b.cache[CacheKey{stock, robots, minutesLeft}]; ok && val > geodes {
		return
	}

	b.cache[CacheKey{stock, robots, minutesLeft}] = geodes
}

func (b Blueprint) Next(stock, robots Material, minutesLeft int) int {
	if len(b.cache)%100000 == 0 {
		fmt.Printf("\r%d", len(b.cache))
	}

	// exit recursion if time has run out
	if minutesLeft <= 0 {
		return stock.geode
	}

	if geodes, ok := b.cache[CacheKey{stock, robots, minutesLeft}]; ok {
		return geodes
	}

	// collect
	stock = stock.Add(robots)

	// divide and conquer on buying robots
	maxGeodes := b.Next(stock, robots, minutesLeft-1)
	b.Cache(stock, robots, maxGeodes, minutesLeft-1)

	for material, cost := range b.materials {
		if !stock.IsAffordable(cost) {
			continue
		}

		geodes := b.Next(stock.Sub(cost), robots.Add(robotMap[material]), minutesLeft-1)
		b.Cache(stock.Sub(cost), robots.Add(robotMap[material]), geodes, minutesLeft-1)
		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	/*if maxGeodes > 9 {
		panic(maxGeodes)
	}*/

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
		result := b.Next(Material{}, Material{ore: 1}, 24)

		sum += (i + 1) * result
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
	fmt.Printf("\n%d\n", part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
