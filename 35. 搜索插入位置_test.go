package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	Run(t, searchInsert, []Test{
		{
			Args{Str2Slice("[1,3,5,6]"), 5},
			Want{2},
		},
		{
			Args{Str2Slice("[1,3,5,6]"), 2},
			Want{1},
		},
		{
			Args{Str2Slice("[1,3,5,6]"), 7},
			Want{4},
		},
		{
			Args{Str2Slice("[1,3,5,6]"), 0},
			Want{0},
		},
	})
}
