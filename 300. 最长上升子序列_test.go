package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	Run(t, lengthOfLIS, []Test{
		{
			Args{Str2Slice("[]")},
			Want{0},
		},
		{
			Args{Str2Slice("[10,9,2,5,3,7,101,18]")},
			Want{4},
		},
		{
			Args{Str2Slice("[2]")},
			Want{1},
		},
	})
}
