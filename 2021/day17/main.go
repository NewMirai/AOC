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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// Solve problem
func Solve(s *string) int {
	// Parsing
	parts := strings.Split(strings.TrimSpace(*s), "=")
	xBounds := strings.Split(strings.Split(parts[1], ",")[0], "..")
	yBounds := strings.Split(parts[2], "..")
	lowX, _ := strconv.Atoi(xBounds[0])
	highX, _ := strconv.Atoi(xBounds[1])
	lowY, _ := strconv.Atoi(yBounds[0])
	highY, _ := strconv.Atoi(yBounds[1])
	// Brute force
	max_y := 0
	counts_velocities := 0
	for y := lowY; y < Abs(lowY); y++ {
		for x := 1; x < highX+1; x++ {
			vx, vy := x, y
			x_position, y_position := 0, 0
			max_y_trajectory := 0
			// 216 par dichotomie...
			// smarter by Theo 2 * abs(lowY) + 1
			for t := 0; t < 2*Abs(lowY)+1; t++ {
				x_position += vx
				y_position += vy
				switch {
				case vx > 0:
					vx--
				case vx < 0:
					vx++
				}
				vy--
				max_y_trajectory = Max(max_y_trajectory, y_position)
				if lowX <= x_position && x_position <= highX && lowY <= y_position && y_position <= highY {
					counts_velocities++
					max_y = Max(max_y, max_y_trajectory)
					break
				}
			}
		}
	}
	// part1 return max_y
	return counts_velocities
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
