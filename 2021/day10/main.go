package main

import (
	"container/list"
	"fmt"
	"os"
	"sort"
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
	scores := make([]int, 0)
	for _, line := range lines {
		corrupted := false
		listChars := list.New()
	loop:
		for _, c := range strings.TrimSpace(line) {
			// Part 1
			switch c {
			case '(', '[', '{', '<':
				listChars.PushBack(c)
			case ')':
				if listChars.Back().Value.(rune) != '(' {
					total += 3
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}
			case ']':
				if listChars.Back().Value.(rune) != '[' {
					total += 57
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}

			case '>':
				if listChars.Back().Value.(rune) != '<' {
					total += 25137
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}

			case '}':
				if listChars.Back().Value.(rune) != '{' {
					total += 1197
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}

			}
		}
		// Part 2
		if !corrupted {
			score := 0
			m := map[rune]int{
				'(': 1,
				'[': 2,
				'{': 3,
				'<': 4,
			}
			for e := listChars.Back(); e != nil; e = e.Prev() {
				value := e.Value.(rune)
				score = score*5 + m[value]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	total = scores[len(scores)/2]
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
