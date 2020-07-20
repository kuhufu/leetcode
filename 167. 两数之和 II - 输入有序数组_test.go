package leetcode

import "testing"

func Test_twoSum(t *testing.T) {
	Run(t, twoSum, []Test{
		{
			Args{Str2Slice("[2, 7, 11, 15]"), 9},
			Want{Str2Slice("[1,2]")},
		},
		{
			Args{Str2Slice("[0,0,3,4]"), 0},
			Want{Str2Slice("[1,2]")},
		},
	})
}
