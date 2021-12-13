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

type Position struct {
	x int
	y int
}

func checkPosition(ps map[Position]int, p Position) bool {
	if _, ok := ps[p]; ok {
		ps[p]++
		return true
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkDiag(x1 int, y1 int, x2 int, y2 int) bool {
	p1 := Abs(x2 - x1)
	p2 := Abs(y2 - y1)
	if p1 == p2 {
		return true
	}
	return false
}

// Solve problem
func Solve(s *string) int {
	var total int
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	Positions := make(map[Position]int)
	for _, line := range lines {
		lines := strings.Split(line, " -> ")
		l1 := strings.Split(lines[0], ",")
		l2 := strings.Split(lines[1], ",")
		x1, _ := strconv.Atoi(l1[0])
		y1, _ := strconv.Atoi(l1[1])
		x2, _ := strconv.Atoi(l2[0])
		y2, _ := strconv.Atoi(l2[1])
		switch {
		case x1 == x2:
			switch {
			case y1 == y2:
				continue
			case y1 > y2:
				for i := y2; i <= y1; i++ {
					p := Position{x: x1, y: i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			case y2 > y1:
				for i := y1; i <= y2; i++ {
					p := Position{x: x1, y: i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			}
		case y1 == y2:
			switch {
			case x1 == x2:
				continue
			case x1 > x2:
				for i := x2; i <= x1; i++ {
					p := Position{x: i, y: y1}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			case x2 > x1:
				for i := x1; i <= x2; i++ {
					p := Position{x: i, y: y1}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			}
		// Part 2
		case checkDiag(x1, y1, x2, y2):
			ic := Abs(x2 - x1)
			switch {
			case x1 < x2 && y1 < y2:
				for i := 0; i <= ic; i++ {
					p := Position{x: x1 + i, y: y1 + i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			case x1 < x2 && y1 > y2:
				for i := 0; i <= ic; i++ {
					p := Position{x: x1 + i, y: y1 - i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			case x1 > x2 && y1 < y2:
				for i := 0; i <= ic; i++ {
					p := Position{x: x1 - i, y: y1 + i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			case x1 > x2 && y1 > y2:
				for i := 0; i <= ic; i++ {
					p := Position{x: x1 - i, y: y1 - i}
					if checkPosition(Positions, p) {
						continue
					} else {
						Positions[p] = 1
					}
				}
			}
		default:
			continue
		}
	}
	for _, v := range Positions {
		if v > 1 {
			total++
		}
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
