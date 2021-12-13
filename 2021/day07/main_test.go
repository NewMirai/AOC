package main

import (
	"testing"
)

func BenchmarkTestDay07(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := ReadInput("input.txt")
		Solve(&input)
	}
}
