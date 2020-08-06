package leetcode

import "testing"

func Test_twoSum(t *testing.T) {
	Run(t, twoSum, []Test{
		{
			Args{Slice("[2, 7, 11, 15]").Ints(), 9},
			Want{Slice("[1,2]").Ints()},
		},
		{
			Args{Slice("[0,0,3,4]").Ints(), 0},
			Want{Slice("[1,2]").Ints()},
		},
	})
}
