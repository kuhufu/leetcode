package leetcode

import (
	"testing"
)

func Test_findMin(t *testing.T) {
	Run(t, findMin, []Test{
		{
			Args{Slice("[1,2]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[2,2,2,2,2,0,1,2]").Ints()},
			Want{0},
		},
		{
			Args{Slice("[2,2,2,0,1]").Ints()},
			Want{0},
		},
		{
			Args{Slice("[1,3,5]").Ints()},
			Want{1},
		},
	})
}
