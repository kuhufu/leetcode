package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	Run(t, maxDepth, []Test{
		{
			Args{Tree("[3,9,20,null,null,15,7]")},
			Want{3},
		},
	})
}
