package leetcode

import "testing"

func Test_findLength(t *testing.T) {
	Run(t, findLength, []Test{
		{
			Args{Slice("[1,2,3,2,1]"), Slice("[3,2,1,4,7]").Ints()},
			Want{3},
		},
	})
}
