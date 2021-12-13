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
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	for i := 0; i < len(numbers); i++ {
		// Part1
		// if i+1 < len(numbers){
		//	if numbers[i+1] > numbers[i]{
		//		total++
		//	}
		if i+3 < len(numbers) {
			sum1 := numbers[i] + numbers[i+1] + numbers[i+2]
			sum2 := numbers[i+1] + numbers[i+2] + numbers[i+3]
			if sum2 > sum1 {
				total++
			}
		}
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
