package leetcode

import (
	"testing"
)

func Test_longestOnes(t *testing.T) {
	Run(t, longestOnes, []Test{
		{
			Args{Slice("[1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]").Ints(), 2},
			Want{6},
		},
		{
			Args{Slice("[0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1]").Ints(), 3},
			Want{10},
		},
	})
}
