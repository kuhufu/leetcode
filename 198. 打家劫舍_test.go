package leetcode

import "testing"

func Test_rob(t *testing.T) {
	Run(t, rob, []Test{
		{
			Args{Str2Slice("[1,2,3,1]")},
			Want{4},
		},
		{
			Args{Str2Slice("[2,1,1,2]")},
			Want{4},
		},

		{
			Args{Str2Slice("[2,7,9,3,1]")},
			Want{12},
		},
		{
			Args{Str2Slice("[1,2]")},
			Want{2},
		},
		{
			Args{Str2Slice("[1]")},
			Want{1},
		},
	})
}
