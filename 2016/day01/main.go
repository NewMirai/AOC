package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadInput of the problem
func ReadInput(f string) string {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type Instruction struct {
	direction string
	unit      int
}

type Position struct {
	x int
	y int
}

type Set map[Position]bool

// Rotate given direction and initial position
func Rotate(d *string, o *int) {
	switch *d {
	case "L":
		*o = (*o + 1) % 4
	case "R":
		*o = (*o + 3) % 4
	}
}

// Move coordinates from desired number of units
func Move(unit *int, p *Position, o *int) {
	switch *o {
	case 0:
		p.y += *unit
	case 1:
		p.x -= *unit
	case 2:
		p.y -= *unit
	case 3:
		p.x += *unit
	}
}

func Abs(x int) int {
	switch {
	case x < 0:
		x = -x
	default:
	}
	return x
}

func ComputeDistance(x *int, y *int) (d int) {
	return Abs(*x) + Abs(*y)
}

// Solve problem PART 1
func Solve(s *string) (res int) {
	instructions := strings.Split(strings.TrimSpace(*s), ", ")
	// Initial position
	var p Position
	var o int
	for _, instruction := range instructions {
		unit, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}
		it := Instruction{string(instruction[0]), unit}
		Rotate(&it.direction, &o)
		Move(&it.unit, &p, &o)
	}
	res = ComputeDistance(&p.x, &p.y)
	return res
}

// Solve problem PART 2
func Solve2(s *string) (res int) {
	instructions := strings.Split(strings.TrimSpace(*s), ", ")
	// Initial position
	var p Position
	var o int
	locations := make(Set)
outer:
	for _, instruction := range instructions {
		unit, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}
		it := Instruction{string(instruction[0]), unit}
		Rotate(&it.direction, &o)
		for i := 1; i <= unit; i++ {
			step := 1
			Move(&step, &p, &o)
			if _, ok := locations[p]; ok {
				break outer
			} else {
				locations[p] = true
			}
		}
	}
	res = ComputeDistance(&p.x, &p.y)
	return res
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
