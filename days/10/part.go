package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type CPU struct {
	X       int
	History []int
}

func NewCPU() CPU {
	c := CPU{}
	c.X = 1
	c.History = make([]int, 0)

	// regarding to the right solution, the first
	// cycle is reserved for adding the initial 1
	c.History = append(c.History, c.X)
	return c
}

func (c *CPU) Noop() {
	c.History = append(c.History, c.X)
}

func (c *CPU) AddX(value int) {
	c.Noop()
	c.X += value
	c.Noop()
}

func (c *CPU) Run(program []string) {
	for _, line := range program {
		parts := strings.Fields(line)

		switch parts[0] {
		case "noop":
			c.Noop()
		case "addx":
			num := s_strings.ToInt(parts[1])
			c.AddX(num)
		default:
			panic(fmt.Sprintf("unknown command %s", parts[0]))
		}
	}
}

func (c CPU) SignalStrength() int {
	cycles := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, cycle := range cycles {
		sum += cycle * c.History[cycle-1]
	}
	return sum
}

func (c CPU) CRT() string {
	var screen string
	for i, value := range c.History {
		// Position rolls over to range [0,39]
		i = i % 40

		// sprite is three pixels wide, so check with a tolerance of +/- 1
		// if the pixel should be drawn
		if math.Within(value, i-1, i+1) {
			screen += "#"
		} else {
			screen += " "
		}

		// add new line every 40 chars
		// add +1 because i is an index and 0%40=0 which would result in
		// a leading new line
		if (i+1)%40 == 0 {
			screen += "\n"
		}
	}

	return screen
}

func part1(data []string) int {
	cpu := NewCPU()
	cpu.Run(data)
	return cpu.SignalStrength()
}

func part2(data []string) string {
	cpu := NewCPU()
	cpu.Run(data)
	return cpu.CRT()
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Print(part2(data))
}
