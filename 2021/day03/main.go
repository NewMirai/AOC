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

// Solve problem part 1
func Solve1(s *string) (total int64) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	Grid := make([][]rune, len(lines))
	// Populate the grid
	for i, line := range lines {
		numbers := make([]rune, len(line))
		for idx, num := range line {
			numbers[idx] = num
		}
		Grid[i] = numbers
	}
	lcol := len(Grid[0])
	lgr := make([]string, lcol)
	for j := 0; j < lcol; j++ {
		var c0 int
		var c1 int
		var mc string
		for i := 0; i < len(lines); i++ {
			el := Grid[i][j]
			if el == '0' {
				c0++
			} else {
				c1++
			}
			if c1 > c0 {
				mc = "1"
			} else {
				mc = "0"
			}
		}
		lgr[j] = mc
	}
	sgr := strings.Join(lgr, "")
	var ser string
	for _, r := range sgr {
		if r == '0' {
			ser += "1"
		} else {
			ser += "0"
		}
	}
	gr, _ := strconv.ParseInt(sgr, 2, 8)
	er, _ := strconv.ParseInt(ser, 2, 8)
	total = gr * er
	return total
}

// Solve problem part 2
func Solve2(s *string) (total int64) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	lines2 := make([]string, len(lines))
	copyN := copy(lines2, lines)
	fmt.Println(copyN)
	Grid := make([][]rune, len(lines))
	for i, line := range lines {
		numbers := make([]rune, len(lines[0]))
		for idx, num := range line {
			numbers[idx] = num
		}
		Grid[i] = numbers
	}
	lcol := len(Grid[0])
	for len(lines) > 1 {
		for j := 0; j < lcol; j++ {
			linesFiltered := make([]string, 0)
			c0 := 0
			c1 := 0
			var mc string
			for i := 0; i < len(lines); i++ {
				if lines[i][j] == '0' {
					c0++
				} else {
					c1++
				}
				if c1 >= c0 {
					mc = "1"
				} else {
					mc = "0"
				}
			}
			for _, line := range lines {
				if string(line[j]) == mc {
					linesFiltered = append(linesFiltered, line)
				}
			}
			// update
			lines = linesFiltered
		}
	}

loop:
	for len(lines2) > 1 {
		for j := 0; j < lcol; j++ {
			linesFiltered2 := make([]string, 0)
			c0 := 0
			c1 := 0
			var lc string
			for i := 0; i < len(lines2); i++ {
				if lines2[i][j] == '0' {
					c0++
				} else {
					c1++
				}
				if c1 >= c0 {
					lc = "1"
				} else {
					lc = "0"
				}
			}
			for _, line := range lines2 {
				if len(lines2) == 1 {
					break loop
				}
				if string(line[j]) != lc {
					linesFiltered2 = append(linesFiltered2, line)
				}
			}
			// update
			lines2 = linesFiltered2
		}
	}
	sgr := lines[0]
	ser := lines2[0]
	gr, _ := strconv.ParseInt(sgr, 2, 64)
	er, _ := strconv.ParseInt(ser, 2, 64)
	total = gr * er
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve1(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
