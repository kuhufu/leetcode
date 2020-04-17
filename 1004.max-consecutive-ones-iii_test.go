package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestOnes(t *testing.T) {
	tests := []struct {
		A   []int
		K   int
		out int
	}{
		{
			A:   []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			K:   2,
			out: 6,
		},
		{
			A:   []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			K:   3,
			out: 10,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.A), func(t *testing.T) {
			out := longestOnes(test.A, test.K)
			if test.out != out {
				t.Errorf("want %v, but got %v", test.out, out)
			}
		})
	}
}