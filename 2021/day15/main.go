package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func MinPathSum(grid [][]int, H int, W int) int {
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for j := 1; j < W; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			dp[i][j] = grid[i][j] + Min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[H-1][W-1] - dp[0][0]
}

func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	H := len(lines)
	W := len(lines[0])
	grid := make([][]int, H)
	for i, line := range lines {
		grid[i] = make([]int, W)
		for j := 0; j < W; j++ {
			cost, _ := strconv.Atoi(string(line[j]))
			grid[i][j] = cost
		}
	}
	total = MinPathSum(grid, H, W)
	return total
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
