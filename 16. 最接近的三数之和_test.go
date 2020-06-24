package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	Run(t, threeSumClosest, []Test{
		{
			Args{Str2Slice("[0,2,1,-3]"), 1},
			Want{0},
		},
		{
			Args{Str2Slice("[1,1,1,0]"), 100},
			Want{3},
		},
		{
			Args{Str2Slice("[-1,2,1,-4]"), 1},
			Want{2},
		},
	})
}
