package leetcode

import "testing"

func Test_treeSum(t *testing.T) {
	Run(t, threeSum, []Test{

		{
			Args{Str2Slice("[-1, 0, 1, 2, -1, -4]")},
			Want{Str22DSlice("[[-1, -1, 2],[-1, 0, 1]]")},
		},
		{
			Args{Str2Slice("[0,0,0,0]")},
			Want{Str22DSlice("[[0,0,0]]")},
		},
	})
}
