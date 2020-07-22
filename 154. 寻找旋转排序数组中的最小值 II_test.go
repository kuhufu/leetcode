package leetcode

import (
	"testing"
)

func Test_findMin(t *testing.T) {
	Run(t, findMin, []Test{
		{
			Args{Str2Slice("[1,2]")},
			Want{1},
		},
		{
			Args{Str2Slice("[2,2,2,2,2,0,1,2]")},
			Want{0},
		},
		{
			Args{Str2Slice("[2,2,2,0,1]")},
			Want{0},
		},
		{
			Args{Str2Slice("[1,3,5]")},
			Want{1},
		},
	})
}
