package leetcode

import "testing"

func Test_generateTrees(t *testing.T) {
	Run(t, generateTrees, []Test{
		{
			Args{3},
			Want{[]*TreeNode{
				Str2Tree("[1,null,2,null,3]"),
				Str2Tree("[1,null,3,2]"),
				Str2Tree("[2,1,3]"),
				Str2Tree("[3,1,null,null,2]"),
				Str2Tree("[3,2,null,1]"),
			},
			},
		},
	})
}
