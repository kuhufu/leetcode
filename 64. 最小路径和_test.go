package leetcode

import "testing"

func Test_minPathSum(t *testing.T) {
	Run(t, minPathSum, []Test{
		{
			Args{Str22DSlice("[[1,3,1],[1,5,1],[4,2,1]]")},
			Want{7},
		},
	})
}
