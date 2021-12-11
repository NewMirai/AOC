package main

import (
	"fmt"
	"log"
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

func (p *Position) GetAdjacentPositions() []Position {
	adj := make([]Position, 8)
	adj[0] = Position{x: p.x, y: p.y - 1}
	adj[1] = Position{x: p.x + 1, y: p.y - 1}
	adj[2] = Position{x: p.x + 1, y: p.y}
	adj[3] = Position{x: p.x + 1, y: p.y + 1}
	adj[4] = Position{x: p.x, y: p.y + 1}
	adj[5] = Position{x: p.x - 1, y: p.y + 1}
	adj[6] = Position{x: p.x - 1, y: p.y}
	adj[7] = Position{x: p.x - 1, y: p.y - 1}
	return adj
}

type Grid struct {
	grid map[Position]*Object
}

func (g *Grid) Display(nrow int, ncol int) {
	for i := 0; i < nrow; i++ {
		row := make([]int, ncol)
		for j := 0; j < ncol; j++ {
			row[j] = g.grid[Position{x: j, y: i}].value
		}
		fmt.Println(row)
	}
}

type Object struct {
	value int
}

// Solve problem
func Solve(s *string) (total int) {
	var numbers []int
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	// for i, line := range lines {
	// 	fmt.Println("Line", i, ": ", line)
	// }
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	for _, number := range numbers {
		total += number
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
