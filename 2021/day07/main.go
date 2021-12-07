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

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func CumSum(n int64) (cumsum int64) {
	var i int64
	for i = 1; i <= n; i++ {
		cumsum += i
	}
	return cumsum
}

// Solve problem
func Solve(s *string) (total int64) {
	positions := strings.Split(strings.TrimSpace(*s), ",")
	n := len(positions)
	numbers := make([]int64, n)
	for i, pos := range positions {
		number, _ := strconv.ParseInt(pos, 10, 64)
		numbers[i] = number
	}
	gridPos := make([][]int64, n)
	for i := 0; i < n; i++ {
		gridPos[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			gridPos[i][j] = CumSum(Abs((numbers[i] - numbers[j])))
		}
	}
	for _, row := range gridPos {
		var currentMin int64
		for _, d := range row {
			currentMin += d
		}
		if total == 0 || currentMin < total {
			total = currentMin
		}
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.test")
	fmt.Println(Solve(&problem_input))
}
