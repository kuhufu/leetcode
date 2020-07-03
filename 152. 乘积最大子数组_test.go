package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {

	Run(t, maxProduct, []Test{
		{
			Args{Str2Slice("[0,2,2,2,0,2,2]")},
			Want{8},
		},
		{
			Args{Str2Slice("[-1,-2,-3,0]")},
			Want{6},
		},
		{
			Args{Str2Slice("[0,2]")},
			Want{2},
		},
		{
			Args{Str2Slice("[0]")},
			Want{0},
		},

		{
			Args{Str2Slice("[2,3,-2,4]")},
			Want{6},
		},

		{
			Args{Str2Slice("[-2,0,-1]")},
			Want{0},
		},
	})
}
