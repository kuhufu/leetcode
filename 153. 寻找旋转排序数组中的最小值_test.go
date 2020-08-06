package leetcode

import "testing"

func Test_findMin1(t *testing.T) {
	Run(t, findMin1, []Test{
		{
			Args{Slice("[1,2]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[4,5,1,2,3]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[2,3,4,5,1]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[2,1]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[3,4,5,1,2]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[4,5,6,7,0,1,2]").Ints()},
			Want{0},
		},
	})
}
