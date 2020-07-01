package leetcode

import "testing"

func Test_findLength(t *testing.T) {
	Run(t, findLength, []Test{
		{
			Args{Str2Slice("[1,2,3,2,1]"), Str2Slice("[3,2,1,4,7]")},
			Want{3},
		},
	})
}
