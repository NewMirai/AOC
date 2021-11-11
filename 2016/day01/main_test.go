package main

import (
	"reflect"
	"testing"
)

func TestDay01Part1(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "R2, L3", want: 5},
		{input: "R2, R2, R2", want: 2},
		{input: "R5, L5, R5, R3", want: 12},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestDay01Part2(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "R8, R4, R4, R8", want: 4},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve2(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
