package leetcode

import "testing"

func Test_findKthLargest(t *testing.T) {
	Run(t, findKthLargest, []Test{

		{
			Args{Slice("[99,99]").Ints(), 1},
			Want{99},
		},
		{
			Args{Slice("[2,1]").Ints(), 2},
			Want{1},
		},
		{
			Args{Slice("[3,2,1,5,6,4]").Ints(), 2},
			Want{5},
		},
		{
			Args{Slice("[3,2,3,1,2,4,5,5,6]").Ints(), 4},
			Want{4},
		},
	})
}
