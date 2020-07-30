package leetcode

import "testing"

func Test_integerBreak(t *testing.T) {
	Run(t, integerBreak, []Test{
		{
			Args{2},
			Want{1},
		},
		{
			Args{10},
			Want{36},
		},
	})
}
