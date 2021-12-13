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

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CostF(n int) int { return n * (n + 1) / 2 }

func Solve(s *string) int {
	positions := strings.Split(strings.TrimSpace(*s), ",")
	n := len(positions)
	numbers := make([]int, n)
	total := math.MaxInt
	for i, pos := range positions {
		number, _ := strconv.Atoi(pos)
		numbers[i] = number
	}
	for _, num := range numbers {
		fuel := 0
		for _, number := range numbers {
			cost := CostF(Abs(num - number))
			fuel += cost
		}
		if fuel < total {
			total = fuel
		}
	}
	return total
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
