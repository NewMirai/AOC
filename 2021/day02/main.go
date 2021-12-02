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
	var hp int
	var dp int
	var aim int
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		d := fields[0]
		unit, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		switch d {
		case "forward":
			hp += unit
			dp += unit * aim
		case "down":
			aim += unit
		case "up":
			aim -= unit
		}
	}
	total = hp * dp
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
