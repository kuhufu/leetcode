package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestOnes(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want int
	}{
		{
			A:    []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			K:    2,
			want: 6,
		},
		{
			A:    []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			K:    3,
			want: 10,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.A), func(t *testing.T) {
			res := longestOnes(test.A, test.K)
			if test.want != res {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
