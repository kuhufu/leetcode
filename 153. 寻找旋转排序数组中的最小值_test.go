package leetcode

import "testing"

func Test_findMin1(t *testing.T) {
	Run(t, findMin1, []Test{
		{
			Args{Str2Slice("[1,2]")},
			Want{1},
		},
		{
			Args{Str2Slice("[4,5,1,2,3]")},
			Want{1},
		},
		{
			Args{Str2Slice("[2,3,4,5,1]")},
			Want{1},
		},
		{
			Args{Str2Slice("[2,1]")},
			Want{1},
		},
		{
			Args{Str2Slice("[3,4,5,1,2]")},
			Want{1},
		},
		{
			Args{Str2Slice("[4,5,6,7,0,1,2]")},
			Want{0},
		},
	})
}
