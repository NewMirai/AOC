package main

import (
	"fmt"
	"math"
	"os"
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
		if _, ok := m[left]; !ok {
			m[left] = right
		}
	}
	step := 0
	for step < 10 {
		pairs := make([]string, 0)
		for i := 1; i < len(polymer); i++ {
			previous := i - 1
			pair := polymer[previous : i+1]
			pairs = append(pairs, pair)

		}
		n := len(pairs)
		for i, pair := range pairs {
			if ie, ok := m[pair]; ok {
				if i == n-1 {
					pairs[i] = string(pair[0]) + ie + string(pair[1])
				} else {
					pairs[i] = string(pair[0]) + ie
				}
			}
		}
		var newPolymer string
		for _, part := range pairs {
			newPolymer += part
		}
		step++
		polymer = newPolymer
	}
	cb := strings.Count(polymer, "B")
	cc := strings.Count(polymer, "C")
	ch := strings.Count(polymer, "H")
	cn := strings.Count(polymer, "N")
	counts := make([]int, 4)
	counts[0] = cb
	counts[1] = cc
	counts[2] = ch
	counts[3] = cn
	min, max := math.MaxInt, math.MinInt
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	total = max - min
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
