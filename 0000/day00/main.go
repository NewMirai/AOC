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
