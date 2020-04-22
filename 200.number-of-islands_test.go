package leetcode

import (
	"fmt"
	"testing"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '1', '0', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '1'},
			},
			3,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.grid), func(t *testing.T) {
			res := numIslands(test.grid)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
