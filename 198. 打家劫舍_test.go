package leetcode

import "testing"

func Test_rob(t *testing.T) {
	Run(t, rob, []Test{
		{
			Args{Slice("[1,2,3,1]").Ints()},
			Want{4},
		},
		{
			Args{Slice("[2,1,1,2]").Ints()},
			Want{4},
		},

		{
			Args{Slice("[2,7,9,3,1]").Ints()},
			Want{12},
		},
		{
			Args{Slice("[1,2]").Ints()},
			Want{2},
		},
		{
			Args{Slice("[1]").Ints()},
			Want{1},
		},
	})
}
