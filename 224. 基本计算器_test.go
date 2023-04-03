package leetcode

import "testing"

func Test_calculate(t *testing.T) {
	Run(t, calculate, []Test{
		{
			Args{"1-(1+3)/2-(2-1)*2"},
			Want{-3},
		},
		{
			Args{"1-(1+3)/2*(3-1)"},
			Want{-3},
		},
		{
			Args{"1-1+2/2"},
			Want{1},
		},

		{
			Args{"1+2*3"},
			Want{7},
		},
		{
			Args{"1+2/2"},
			Want{2},
		},
		{
			Args{"1+2"},
			Want{3},
		},
		{
			Args{"1-(1+3)/2"},
			Want{-1},
		},
	})
}
