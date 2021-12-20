package main

import (
	"testing"
)

var result int

func BenchmarkDay17(b *testing.B) {
	var r int
	problem_input := ReadInput("input.txt")
	for n := 0; n < b.N; n++ {
		r = Solve(&problem_input)
	}
	result = r
}
