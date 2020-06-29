package leetcode

import "testing"

func Test_findKthLargest(t *testing.T) {
	Run(t, findKthLargest, []Test{

		{
			Args{Str2Slice("[99,99]"), 1},
			Want{99},
		},
		{
			Args{Str2Slice("[2,1]"), 2},
			Want{1},
		},
		{
			Args{Str2Slice("[3,2,1,5,6,4]"), 2},
			Want{5},
		},
		{
			Args{Str2Slice("[3,2,3,1,2,4,5,5,6]"), 4},
			Want{4},
		},
	})
}
