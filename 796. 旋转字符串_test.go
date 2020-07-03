package leetcode

import (
	"testing"
)

func Test_rotateString(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want bool
	}{
		{
			"aa",
			"a",
			false,
		},
		{
			"",
			"",
			true,
		},
		{
			"abcde",
			"cdeab",
			true,
		},
		{
			"abcde",
			"abced",
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.A, test.B), func(t *testing.T) {
			res := rotateString(test.A, test.B)
			if test.want != res {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}

func Test2_rotateString(t *testing.T) {
	Run(t, rotateString, []Test{
		{
			Args{
				"aa",
				"a",
			},
			Want{false},
		},
		{
			Args{
				"",
				"",
			},
			Want{true},
		},
		{
			Args{
				"abcde",
				"cdeab",
			},
			Want{true},
		},
		{
			Args{
				"abcde",
				"abced",
			},
			Want{false},
		},
	})
}
