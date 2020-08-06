package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {

	Run(t, maxProduct, []Test{
		{
			Args{Slice("[0,2,2,2,0,2,2]").Ints()},
			Want{8},
		},
		{
			Args{Slice("[-1,-2,-3,0]").Ints()},
			Want{6},
		},
		{
			Args{Slice("[0,2]").Ints()},
			Want{2},
		},
		{
			Args{Slice("[0]").Ints()},
			Want{0},
		},

		{
			Args{Slice("[2,3,-2,4]").Ints()},
			Want{6},
		},

		{
			Args{Slice("[-2,0,-1]").Ints()},
			Want{0},
		},
	})
}
