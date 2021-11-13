package main

import (
	"reflect"
	"testing"
)

func TestDay03Part1(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "5 10 25", want: 0},
		{input: "5 10 25\n19 7 25", want: 1},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestDay03Part2(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "5 10 25\n5 19 3\n 1 4 8", want: 1},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve2(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
