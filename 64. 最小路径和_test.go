package leetcode

import "testing"

func Test_minPathSum(t *testing.T) {
	Run(t, minPathSum, []Test{
		{
			Args{Slice("[[1,3,1],[1,5,1],[4,2,1]]").DInts()},
			Want{7},
		},
	})
}
