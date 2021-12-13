package main

import (
	"testing"
)

var result int

func BenchmarkDay10(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Solve("input.txt")
	}
	result = r
}
