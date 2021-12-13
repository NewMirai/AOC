package main

import (
	"fmt"
	"math"
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

func (fm FoldMap) FoldAlongX(x int, maxRow int, maxCol int, total *int) {
	tX := maxCol - x + 1
	for p := range fm {
		if p.x >= x {
			newPos := Position{x: p.x - tX, y: p.y}
			if _, ok := fm[newPos]; !ok && p.x-tX >= 0 {
				fm[newPos] = struct{}{}
			} else if _, ok := fm[newPos]; ok {
				delete(fm, newPos)
			}
			delete(fm, p)
		}
	}
	*total = len(fm)
}

func (fm FoldMap) FoldAlongY(y int, maxRow int, maxCol int, total *int) {
	tY := maxRow - y + 1
	for p := range fm {
		if p.y >= y {
			newPos := Position{x: p.x, y: p.y - tY}
			if _, ok := fm[newPos]; !ok && p.y-tY >= 0 {
				fm[newPos] = struct{}{}
			} else if _, ok := fm[newPos]; ok {
				delete(fm, newPos)
			}
			delete(fm, p)
		}
	}
	*total = len(fm)
}

func Solve(s *string) (total int) {
	instructions := strings.Split(strings.TrimSpace(*s), "\n\n")
	coordinates := strings.Split(instructions[0], "\n")
	foldsInstructions := strings.Split(instructions[1], "\n")
	m := make(FoldMap)
	folds := make([]Fold, len(foldsInstructions))
	maxRow, maxCol := math.MinInt, math.MinInt
	for _, line := range coordinates {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		if x > maxCol {
			maxCol = x
		}
		if y > maxRow {
			maxRow = y
		}
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
	for _, f := range folds {
		switch f.axe {
		case "x":
			m.FoldAlongX(f.c, maxRow, maxCol, &total)
		case "y":
			m.FoldAlongY(f.c, maxRow, maxCol, &total)
		}
		fmt.Println(total)
	}
	return total
}

func main() {
	problem_input := ReadInput("input.test")
	fmt.Println(Solve(&problem_input))
}
