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

// Heap Perm

func Perm(k int, a []int) {
	if k == 1 {
		fmt.Println(a)
	}
	Perm(k-1, a)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		Perm(k-1, a)
	}
}

// Solve problem
func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	for _, line := range lines {
		data := strings.Split(line, "|")
		signalPatterns := strings.Fields(data[0])
		outputValues := strings.Fields(data[1])
		fmt.Println(signalPatterns)
		fmt.Println(outputValues)
		digits := map[int]string{
			0: "abcefg",
			1: "cf",
			2: "acdeg",
			3: "acdfg",
			4: "bcdf",
			5: "abdfg",
			6: "abdefg",
			7: "acf",
			8: "abcdefg",
			9: "abcdfg",
		}
		Perm(5, []int{1, 2, 3, 4, 5})
		fmt.Println(digits)
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
	}
	return total
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
