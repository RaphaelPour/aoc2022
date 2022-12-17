package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

var (
	rocks = []Rock{
		Rock{"-", []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, Point{0, 0}, Point{3, 0}},
		Rock{"+", []Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, Point{0, 0}, Point{2, 2}},
		Rock{"L", []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, Point{0, 0}, Point{2, 2}},
		Rock{"|", []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, Point{0, 0}, Point{0, 3}},
		Rock{"X", []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, Point{0, 0}, Point{1, 1}},
	}

	directions = map[byte]Point{
		'<': Point{-1, 0},
		'>': Point{1, 0},
	}
)

type Point struct {
	x, y int
}

func (p *Point) Move(other Point) {
	p.x += other.x
	p.y += other.y
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

type Rock struct {
	name     string
	shape    []Point
	min, max Point
}

type Instance struct {
	rock           *Rock
	transformation Point
}

func (inst *Instance) Transform(t Point, blocked map[Point]struct{}) bool {
	// transform if no collision takes place
	newTransformation := inst.transformation.Add(t)
	destMin := inst.rock.min.Add(newTransformation)
	destMax := inst.rock.max.Add(newTransformation)

	// rock should be in bounds
	if destMin.y < 0 || destMin.x < 0 || destMax.x >= 7 {
		return false
	}

	for _, p := range inst.rock.shape {
		if _, ok := blocked[p.Add(newTransformation)]; ok {
			return false
		}
	}

	inst.transformation = newTransformation
	return true
}

func (inst Instance) Top() int {
	// add one to convert index to actual height where
	// '-' with max-y 0 has a thickness of 1
	return inst.rock.max.Add(inst.transformation).y + 1
}

func (inst Instance) Points() []Point {
	points := make([]Point, len(inst.rock.shape))
	for i := range inst.rock.shape {
		points[i] = inst.rock.shape[i].Add(inst.transformation)
	}
	return points
}

func Dump(blocked map[Point]struct{}) {
	for y := 20; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			if _, ok := blocked[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}

	fmt.Println("+-------+")
}

func part1(data string) int {
	// track points that are blocked by rock for collission detection
	blocked := make(map[Point]struct{})

	// dummy map for setting each new rock to its start point where
	// by definition is nothing blocked
	dummyMap := make(map[Point]struct{})

	// track the rock-instances we already placed
	// TODO: eliminate, we just need to track the height of the heighest instance
	instances := make([]Instance, 0)

	// direction index to cycle through the input
	dir := 0
	for i := 0; len(instances) < 2022; i = (i + 1) % len(rocks) {
		instance := Instance{&rocks[i], Point{0, 0}}

		// floor is zero or the tip of the heighest rock
		floor := 0
		if len(instances) > 0 {
			floor = instances[len(instances)-1].Top()
		}

		// move rock to its start position, it can't be blocked there
		instance.Transform(Point{2, 3 + floor}, dummyMap)
		// fmt.Println("start with height:", 3+floor)
		// move rock until stuck
		for {
			// apply jet stream
			instance.Transform(directions[data[dir]], blocked)
			dir = (dir + 1) % len(data)

			// move downwards
			if !instance.Transform(Point{0, -1}, blocked) {
				break
			}

		}

		for _, p := range instance.Points() {
			blocked[p] = struct{}{}
		}
		// fmt.Println(blocked)
		// Dump(blocked)

		instances = append(instances, instance)
	}

	// fmt.Println(instances[0], instances[0].rock)
	// fmt.Println(blocked)

	Dump(blocked)

	max := 0
	for p := range blocked {
		if p.y > max {
			max = p.y
		}
	}
	fmt.Println(max)

	return instances[len(instances)-1].Top()
}

func part2(data string) int {
	return 0
}

func main() {
	data := input.LoadString("input1")[0]

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	// fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part2(data))
}
