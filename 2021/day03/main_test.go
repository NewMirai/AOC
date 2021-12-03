package main

import (
	"reflect"
	"testing"
)

func TestDay00(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "1\n2\n", want: 3},
		{input: "0\n2\n", want: 2},
		{input: "3\n2\n", want: 5},
		{input: "1\n\n", want: 1},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
