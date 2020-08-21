package leetcode

import "testing"

func Test_minDepth(t *testing.T) {
	Run(t, minDepth, []Test{
		{
			Args{Tree("[1,2]")},
			Want{2},
		},
		{
			Args{Tree("[3,9,20,null,null,15,7]")},
			Want{2},
		},
	})
}
