package main

import (
	"reflect"
	"testing"
)

func TestDay02(t *testing.T) {
	type TableTests struct {
		input string
		want  int
	}

	// Write your test cases
	tests := []TableTests{
		{input: "forward 5\ndown 5\nforward 8\nup 3\n down 8\n forward 2", want: 900},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
