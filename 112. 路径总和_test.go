package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	Run(t, hasPathSum, []Test{
		{
			Args{Str2Tree("[5,4,8,11,null,13,4,7,2,null,null,null,1]"), 22},
			Want{true},
		},
	})
}
