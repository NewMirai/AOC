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
	numbers := make(map[int]int)
	numbersCopy := make(map[int]int)
	lines := strings.Split(strings.TrimSpace(*s), ",")
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := numbers[number]; ok {
			numbers[number]++
		} else {
			numbers[number] = 1
		}
	}
	day := 0
	for day < 3 {
		// Make a Copy
		for k, v := range numbers {
			numbersCopy[k] = v
		}
		for k, v := range numbersCopy {
			switch {
			case k == 0:
				if v > 0 {
					numbers[0]--
					if _, ok := numbers[8]; ok {
						numbers[8]++
					} else {
						numbers[8] = 1
					}
					if _, ok := numbers[6]; ok {
						numbers[6]++
					} else {
						numbers[6] = 1
					}
				}
			default:
				for i := 0; i < v; i++ {
					numbers[k]--
					nk := k - 1
					if nk >= 0 {
						if _, ok := numbers[nk]; ok {
							numbers[nk]++
						} else {
							numbers[nk] = 1
						}
					}
				}
			}
		}
		day++
		fmt.Println("After day", day)
		fmt.Println(numbers)
	}
	for _, v := range numbers {
		total += v
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.test")
	fmt.Println(Solve(&problem_input))
}
