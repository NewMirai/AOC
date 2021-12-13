package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Solve(s string) (total int) {
	// Read line by line
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scores := make([]int, 0)
	listChars := list.New()
	for scanner.Scan() {
		corrupted := false
	loop:
		for _, c := range strings.TrimSpace(scanner.Text()) {
			// Part 1
			switch c {
			case '(', '[', '{', '<':
				listChars.PushBack(c)
			case ')':
				if listChars.Back().Value.(rune) != c-1 {
					total += 3
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}
			case ']':
				if listChars.Back().Value.(rune) != c-2 {
					total += 57
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}

			case '>':
				if listChars.Back().Value.(rune) != c-2 {
					total += 25137
					corrupted = true
					break loop
				} else {
					listChars.Remove(listChars.Back())
				}

			case '}':
				if listChars.Back().Value.(rune) != c-2 {
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
				score = score*5 + m[e.Value.(rune)]
			}
			scores = append(scores, score)
		}
		listChars.Init()
	}
	sort.Ints(scores)
	total = scores[len(scores)/2]
	return
}

// main function
func main() {
	fmt.Println(Solve("input.txt"))
}
