package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	Run(t, searchInsert, []Test{
		{
			Args{Slice("[1,3,5,6]").Ints(), 5},
			Want{2},
		},
		{
			Args{Slice("[1,3,5,6]").Ints(), 2},
			Want{1},
		},
		{
			Args{Slice("[1,3,5,6]").Ints(), 7},
			Want{4},
		},
		{
			Args{Slice("[1,3,5,6]").Ints(), 0},
			Want{0},
		},
	})
}
