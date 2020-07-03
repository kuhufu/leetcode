package leetcode

import (
	"testing"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			Str22DSlice("[[1]]"),
			-1,
		},
		{
			Str22DSlice("[[0]]"),
			1,
		},
		{
			Str22DSlice("[[1,0,0],[1,1,0],[1,1,0]]"),
			-1,
		},

		{
			Str22DSlice("[[0,0,0],[1,1,0],[1,1,0]]"),
			4,
		},
		{
			Str22DSlice("[[0,1],[1,0]]"),
			2,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.grid), func(t *testing.T) {
			res := shortestPathBinaryMatrix(test.grid)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
