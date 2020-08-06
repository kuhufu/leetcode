package leetcode

import (
	"testing"
)

func Test_generate(t *testing.T) {
	Run(t, generate, []Test{
		{
			Args{5},
			Want{[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			}},
		},
	})
}
