package leetcode

import "testing"

func Test_binarySearch(t *testing.T) {
	Run(t, binarySearch, []Test{
		{
			Args{Str2Slice("[5]"), 6},
			Want{0},
		},
		{
			Args{Str2Slice("[-1,0,3,5,9,12]"), 9},
			Want{4},
		},
		{
			Args{Str2Slice("[-1,0,3,5,9,12]"), 2},
			Want{-1},
		},
	})
}

func Test_golangStdBinarySearch(t *testing.T) {
	Run(t, golangStdBinarySearch2, []Test{
		{
			Args{Str2Slice("[5]"), 5},
			Want{0},
		},
		{
			Args{Str2Slice("[-1,0,3,5,9,12]"), 9},
			Want{4},
		},
		{
			Args{Str2Slice("[-1,0,3,5,9,12]"), 15},
			Want{6},
		},
		{
			Args{Str2Slice("[-1,0,3,5,9,12]"), 2},
			Want{2},
		},
	})
}
