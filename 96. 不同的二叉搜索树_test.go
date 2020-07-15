package leetcode

import "testing"

func Test_numTrees(t *testing.T) {
	Run(t, numTrees, []Test{
		{
			Args{3},
			Want{5},
		},
		{
			Args{4},
			Want{14},
		},
	})
}
