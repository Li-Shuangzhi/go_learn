package _struct

import "testing"

func TestAbs(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{-1, 1},
		{-11, 11},
		{-12, 12},
		{0, 0},
		{1, 1},
	}
	for _, test := range tests {
		if got := Abs(test.input); got != test.want {
			t.Errorf("Abs(%d) = %d", test.input, got)
		}
	}
}
