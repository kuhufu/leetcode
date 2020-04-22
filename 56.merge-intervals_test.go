package leetcode

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{2, 3}, {1, 4}},
			[][]int{{1, 4}},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.intervals), func(t *testing.T) {
			res := merge(test.intervals)

			if !Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}