package leetcode

import "testing"

func Test_binarySearch(t *testing.T) {
	Run(t, binarySearch, []Test{
		{
			Args{Slice("[5]").Ints(), 6},
			Want{0},
		},
		{
			Args{Slice("[-1,0,3,5,9,12]").Ints(), 9},
			Want{4},
		},
		{
			Args{Slice("[-1,0,3,5,9,12]").Ints(), 2},
			Want{-1},
		},
	})
}

func Test_golangStdBinarySearch(t *testing.T) {
	Run(t, golangStdBinarySearch2, []Test{
		{
			Args{Slice("[5]").Ints(), 5},
			Want{0},
		},
		{
			Args{Slice("[-1,0,3,5,9,12]").Ints(), 9},
			Want{4},
		},
		{
			Args{Slice("[-1,0,3,5,9,12]").Ints(), 15},
			Want{6},
		},
		{
			Args{Slice("[-1,0,3,5,9,12]").Ints(), 2},
			Want{2},
		},
	})
}
