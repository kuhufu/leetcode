package leetcode

import "testing"

func Test_longestMountain(t *testing.T) {
	Run(t, longestMountain, []Test{
		{
			Args{Slice("[2,1,4,7,3,2,5,2,1]").Ints()},
			Want{5},
		},
		{
			Args{Slice("[2,1,4,7,3,2,5]").Ints()},
			Want{5},
		},
		{
			Args{Slice("[2,2,2]").Ints()},
			Want{0},
		},
	})
}
