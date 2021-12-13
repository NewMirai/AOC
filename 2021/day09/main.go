package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput of the problem
func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type Position struct {
	x int
	y int
}

func (p *Position) GetNeighbours() []Position {
	arr := make([]Position, 0)
	arr = append(arr, Position{x: p.x + 1, y: p.y})
	arr = append(arr, Position{x: p.x, y: p.y + 1})
	arr = append(arr, Position{x: p.x - 1, y: p.y})
	arr = append(arr, Position{x: p.x, y: p.y - 1})
	return arr
}

func Crawl(m *map[int][]Position, p *Position, idx int) {
}

func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	positions := make(map[Position]int)
	for y, line := range lines {
		for x, v := range line {
			height, _ := strconv.Atoi(string(v))
			positions[Position{x: x, y: y}] = height
		}
	}
	lowPoints := make([]Position, 0)
	for p := range positions {
		neighbours := p.GetNeighbours()
		height := positions[p]
		n := 0
		count := 0
		for _, neighbourPos := range neighbours {
			if _, isValid := positions[neighbourPos]; isValid {
				neighboursHeight := positions[neighbourPos]
				n++
				switch {
				case neighboursHeight-height > 0:
					count++
				}
			}
		}
		if n == count {
			lowPoints = append(lowPoints, p)
		}
	}
	// Part 1
	for _, lp := range lowPoints {
		total += positions[lp] + 1
	}
	// Part 2
	for _, lp := range lowPoints {
		// A basin - get all neighbours
		// valid != 9 && above current value
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
