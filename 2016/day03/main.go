package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

// IsTriangle check is three numbers gives a IsTriangle
func IsTriangle(numbers []int) (result bool) {
	sort.Ints(numbers)
	if numbers[0]+numbers[1] > numbers[2] {
		result = true
	} else {
		result = false
	}
	return result
}

// Solve problem Part 1
func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	for _, line := range lines {
		s := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(line), 3)
		var numbers []int
		for _, n := range s {
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)

		}
		if IsTriangle(numbers) {
			total++
		}
	}
	return total
}

// Solve problem Part 2
func Solve2(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	Grid := make([][]int, len(lines))
	for i, line := range lines {
		s := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(line), 3)
		var numbers []int
		for _, n := range s {
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)

		}
		Grid[i] = numbers
	}
	for j := 0; j < 3; j++ {
		var sides []int
		for i := 0; i < len(lines); i++ {
			side := Grid[i][j]
			sides = append(sides, side)
			if len(sides) == 3 {
				if IsTriangle(sides) {
					total++
				}
				// Clear the sides
				sides = nil
			}
		}
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
