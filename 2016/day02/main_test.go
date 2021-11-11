package main

import (
	"reflect"
	"testing"
)

func TestDay00(t *testing.T) {
	type TableTests struct {
		input string
		want  string
	}

	// Write your test cases
	tests := []TableTests{
		{input: "ULL\nRRDDD\nLURDL\nUUUUD", want: "1985"},
	}

	for _, tc := range tests {
		problem_input := tc.input
		got := Solve(&problem_input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

}
