package leetcode

import (
	"testing"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	Run(t, shortestPathBinaryMatrix, []Test{
		{
			Args{Slice("[[1]]").DInts()},
			Want{-1},
		},
		{
			Args{Slice("[[0]]").DInts()},
			Want{1},
		},
		{
			Args{Slice("[[1,0,0],[1,1,0],[1,1,0]]").DInts()},
			Want{-1},
		},

		{
			Args{Slice("[[0,0,0],[1,1,0],[1,1,0]]").DInts()},
			Want{4},
		},
		{
			Args{Slice("[[0,1],[1,0]]").DInts()},
			Want{2},
		},
	})
}
