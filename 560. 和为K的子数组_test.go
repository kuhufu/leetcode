package leetcode

import "testing"

func Test_subarraySum(t *testing.T) {
	Run(t, subarraySum, []Test{
		{
			Args{
				Str2Slice("[-1,-1,1]"),
				0,
			},
			Want{1},
		},
		{
			Args{
				Str2Slice("[1]"),
				0,
			},
			Want{0},
		},
		{
			Args{
				Str2Slice("[1,1,2,1]"),
				2,
			},
			Want{2},
		},
		{
			Args{
				Str2Slice("[1,1,2,1]"),
				3,
			},
			Want{2},
		},
		{
			Args{
				Str2Slice("[1,1,1]"),
				2,
			},
			Want{2},
		},
	})
}
