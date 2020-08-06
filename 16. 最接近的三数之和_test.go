package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	Run(t, threeSumClosest, []Test{
		{
			Args{Slice("[0,2,1,-3]").Ints(), 1},
			Want{0},
		},
		{
			Args{Slice("[1,1,1,0]").Ints(), 100},
			Want{3},
		},
		{
			Args{Slice("[-1,2,1,-4]").Ints(), 1},
			Want{2},
		},
	})
}
