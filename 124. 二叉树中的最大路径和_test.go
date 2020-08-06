package leetcode

import "testing"

func Test_maxPathSum(t *testing.T) {
	Run(t, maxPathSum, []Test{
		{
			Args{Tree("[1,-2,3]").Parse()},
			Want{4},
		},
		{
			Args{Tree("[-2,null,-3]").Parse()},
			Want{-2},
		},
		{
			Args{Tree("[-3]").Parse()},
			Want{-3},
		},
		{
			Args{Tree("[-10,9,20,null,null,15,7]").Parse()},
			Want{42},
		},
		{
			Args{Tree("[1,2,3]").Parse()},
			Want{6},
		},
	})
}
