package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	Run(t, lengthOfLIS, []Test{
		{
			Args{Slice("[]").Ints()},
			Want{0},
		},
		{
			Args{Slice("[10,9,2,5,3,7,101,18]").Ints()},
			Want{4},
		},
		{
			Args{Slice("[2]").Ints()},
			Want{1},
		},
	})
}
