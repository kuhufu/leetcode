package leetcode

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.numRows), func(t *testing.T) {
			res := generate(test.numRows)

			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
