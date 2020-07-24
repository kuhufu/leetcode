package leetcode

import "testing"

func Test_divisorGame(t *testing.T) {
	Run(t, divisorGame, []Test{
		{
			Args{1},
			Want{false},
		},
		{
			Args{2},
			Want{true},
		},
	})
}
