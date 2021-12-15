package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

func Solve(s *string) (total int) {
	parts := strings.Split(strings.TrimSpace(*s), "\n\n")
	polymer := parts[0]
	pairsRules := parts[1]
	m := make(map[string]string)
	for _, pairRule := range strings.Split(pairsRules, "\n") {
		rule := strings.Split(pairRule, " -> ")
		left, right := rule[0], rule[1]
		m[left] = right
	}
	pairsCounter := make(map[string]int)
	for i := 1; i < len(polymer); i++ {
		previous := i - 1
		pair := polymer[previous : i+1]
		pairsCounter[pair]++

	}
	step := 0
	charCounter := make(map[string]int)
	// Part 1
	//for step < 10 {
	for step < 40 {
		updateCounter := make(map[string]int)
		updateCharCounter := make(map[string]int)
		for k, v := range pairsCounter {
			k1, k2 := string(k[0])+m[k], m[k]+string(k[1])
			updateCounter[k1] += v
			updateCounter[k2] += v

			updateCharCounter[string(k[0])] += v
			updateCharCounter[m[k]] += v
		}
		pairsCounter = updateCounter
		charCounter = updateCharCounter
		step++
	}
	charCounter[string(polymer[len(polymer)-1])]++
	counts := make([]int, 0)
	for _, v := range charCounter {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
