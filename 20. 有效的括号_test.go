package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	Run(t, isValid, []Test{
		{
			Args{"("},
			Want{false},
		},
		{
			Args{"()"},
			Want{true},
		},
		{
			Args{"()[]{}"},
			Want{true},
		},
		{
			Args{"([)]"},
			Want{false},
		},
		{
			Args{"{[]}"},
			Want{true},
		},
	})
}
