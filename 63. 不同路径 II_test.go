package leetcode

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	Run(t, uniquePathsWithObstacles, []Test{
		{
			Args{Slice("[[0,0],[1,0]]").DInts()},
			Want{1},
		},
		{
			Args{Slice("[[0,0,0],[0,1,0],[0,0,0]]").DInts()},
			Want{2},
		},

		{
			Args{Slice("[[1]]").DInts()},
			Want{0},
		},
		{
			Args{Slice("[[0,1]]").DInts()},
			Want{0},
		},
	})
}
