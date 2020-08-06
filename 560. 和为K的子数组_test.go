package leetcode

import "testing"

func Test_subarraySum(t *testing.T) {
	Run(t, subarraySum, []Test{
		{
			Args{
				Slice("[-1,-1,1]").Ints(),
				0,
			},
			Want{1},
		},
		{
			Args{
				Slice("[1]").Ints(),
				0,
			},
			Want{0},
		},
		{
			Args{
				Slice("[1,1,2,1]").Ints(),
				2,
			},
			Want{2},
		},
		{
			Args{
				Slice("[1,1,2,1]").Ints(),
				3,
			},
			Want{2},
		},
		{
			Args{
				Slice("[1,1,1]").Ints(),
				2,
			},
			Want{2},
		},
	})
}
