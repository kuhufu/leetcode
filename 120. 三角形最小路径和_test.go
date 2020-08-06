package leetcode

import "testing"

func Test_minimumTotal(t *testing.T) {
	Run(t, minimumTotal, []Test{
		{
			Args{Slice("[[2],[3,4],[6,5,7],[4,1,8,3]]").DInts()},
			Want{11},
		},
	})
}
