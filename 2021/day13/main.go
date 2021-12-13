package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type Position struct {
	x int
	y int
}

type Fold struct {
	axe string
	c   int
}

type FoldMap map[Position]struct{}

func (fm FoldMap) FoldAlongX(x int, total *int) {
	for p := range fm {
		if p.x > x {
			newVal := 2*x - p.x
			newPos := Position{x: newVal, y: p.y}
			if _, ok := fm[newPos]; !ok && newVal >= 0 {
				fm[newPos] = struct{}{}
			}
			delete(fm, p)
		}
	}
	*total = len(fm)
}

func (fm FoldMap) FoldAlongY(y int, total *int) {
	for p := range fm {
		if p.y > y {
			newVal := 2*y - p.y
			newPos := Position{x: p.x, y: newVal}
			if _, ok := fm[newPos]; !ok && newVal >= 0 {
				fm[newPos] = struct{}{}
			}
			delete(fm, p)
		}
	}
	*total = len(fm)
}

func (fm FoldMap) Display() {
	var maxX, maxY int
	for p := range fm {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			pos := Position{x: x, y: y}
			if _, ok := fm[pos]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if x == maxX {
				fmt.Print("\n")
			}
		}
	}
}

func Solve(s *string) (total int) {
	instructions := strings.Split(strings.TrimSpace(*s), "\n\n")
	coordinates := strings.Split(instructions[0], "\n")
	foldsInstructions := strings.Split(instructions[1], "\n")
	m := make(FoldMap)
	folds := make([]Fold, len(foldsInstructions))
	for _, line := range coordinates {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		p := Position{x: x, y: y}
		if _, ok := m[p]; !ok {
			m[p] = struct{}{}
		}
	}
	for i, line := range foldsInstructions {
		parts := strings.Split(strings.Split(line, " ")[2], "=")
		axe := parts[0]
		c, _ := strconv.Atoi(parts[1])
		f := Fold{axe: axe, c: c}
		folds[i] = f
	}
	c := 0
	for _, f := range folds {
		switch f.axe {
		case "x":
			m.FoldAlongX(f.c, &total)
		case "y":
			m.FoldAlongY(f.c, &total)
		}
		// Part 1
		if c == 0 {
			fmt.Println(total)
		}
		c++
	}
	m.Display()
	return total
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
