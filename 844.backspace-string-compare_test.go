package leetcode

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		S    string
		T    string
		want bool
	}{
		{
			"y#fo##f",
			"y#f#o##f",
			true,
		},

		{
			"ab#c",
			"ad#c",
			true,
		},
		{
			"a#c",
			"b",
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.S, test.T), func(t *testing.T) {
			res := backspaceCompare(test.S, test.T)
			if test.want != res {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
