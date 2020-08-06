package leetcode

import "testing"

func Test_treeSum(t *testing.T) {
	Run(t, threeSum, []Test{

		{
			Args{Slice("[-1, 0, 1, 2, -1, -4]").Ints()},
			Want{Slice("[[-1, -1, 2],[-1, 0, 1]]").DInts()},
		},
		{
			Args{Slice("[0,0,0,0]").Ints()},
			Want{Slice("[[0,0,0]]").DInts()},
		},
	})
}
