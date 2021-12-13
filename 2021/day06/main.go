package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// ReadInput of the problem
func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

// Solve problem
func Solve(s *string) (total int) {
	numbers := make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		numbers[i] = 0
	}
	lines := strings.Split(strings.TrimSpace(*s), ",")
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers[number]++
	}
	day := 0
	for day < 256 {
		n0 := numbers[0]
		for i := 0; i < 8; i++ {
			numbers[i] = numbers[i+1]
		}
		numbers[6] += n0
		numbers[8] = n0
		day++
	}
	for _, v := range numbers {
		total += v
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	start := time.Now()
	fmt.Println(Solve(&problem_input))
	elapsed := time.Since(start)
	log.Printf("Solution took: %s", elapsed)
}
