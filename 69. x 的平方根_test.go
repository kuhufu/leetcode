package leetcode

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	Run(t, mySqrt, []Test{
		{
			Args{4},
			Want{2},
		},

		{
			Args{5},
			Want{2},
		},
		{
			Args{10},
			Want{3},
		},
	})
}
