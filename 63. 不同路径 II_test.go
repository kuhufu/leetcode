package leetcode

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	Run(t, uniquePathsWithObstacles, []Test{
		{
			Args{Str22DSlice("[[0,0],[1,0]]")},
			Want{1},
		},
		{
			Args{Str22DSlice("[[0,0,0],[0,1,0],[0,0,0]]")},
			Want{2},
		},

		{
			Args{Str22DSlice("[[1]]")},
			Want{0},
		},
		{
			Args{Str22DSlice("[[0,1]]")},
			Want{0},
		},
	})
}
