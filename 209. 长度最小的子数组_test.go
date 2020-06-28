package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	Run(t, minSubArrayLen, []Test{
		{
			Args{3, Str2Slice("[1,1]")},
			Want{0},
		},
		{
			Args{15, Str2Slice("[5,1,3,5,10,7,4,9,2,8]")},
			Want{2},
		},
		{
			Args{7, Str2Slice("[2,3,1,2,4,3]")},
			Want{2},
		},
	})
}
