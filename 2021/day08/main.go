package main

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput of the problem
func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

// Solve problem
func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	for _, line := range lines {
		data := strings.Split(line, "|")
		signalPatterns := strings.Fields(data[0])
		outputValues := strings.Fields(data[1])
		//Part 1
		// for _, op := range outputValues {
		// 	switch len(op) {
		// 	case 2, 3, 4, 7:
		// 		total++
		// 	default:
		// 		continue
		// 	}
		// }

		//Part 2
		// m := make(map[string]string)
		// for _, sp := range signalPatterns{
		// 	switch len(sp):
		// 	case 2:
		// 	m[
		// }
		// n, _ := strconv.Atoi(number)
		// total += n
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.test")
	fmt.Println(Solve(&problem_input))
}
