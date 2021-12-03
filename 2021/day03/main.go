package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

func RemoveIndex(s [][]rune, index int) [][]rune {
	return append(s[:index], s[index+1:]...)
}

// Solve problem part 1
func Solve1(s *string) (total int64) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	Grid := make([][]rune, len(lines))
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
		c0 := 0
		c1 := 0
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
	gr, _ := strconv.ParseInt(sgr, 2, 64)
	er, _ := strconv.ParseInt(ser, 2, 64)
	total = gr * er
	return total
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// Solve problem part 2
func Solve2(s *string) (total int64) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	Grid := make([][]rune, len(lines))
	numbers := make([]rune, len(lines[0]))
	for i, line := range lines {
		for idx, num := range line {
			numbers[idx] = num
		}
		Grid[i] = numbers
	}
	lcol := len(Grid[0])
	for len(lines) > 1 {
		for j := 0; j < lcol; j++ {
			c0 := 0
			c1 := 0
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
			for i, line := range lines {
				if string(line[j]) == mc {
					runtime.Breakpoint()
					lines = remove(lines, i)
				}
			}
		}
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.test")
	fmt.Println(Solve1(&problem_input))
	fmt.Println(Solve2(&problem_input))
}
