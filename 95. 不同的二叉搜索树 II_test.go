package leetcode

import (
	"testing"
)

func Test_generateTrees(t *testing.T) {
	Run(t, generateTrees, []Test{
		{
			Args{3},
			Want{[]*TreeNode{
				Tree("[1,null,2,null,3]").Parse(),
				Tree("[1,null,3,2]").Parse(),
				Tree("[2,1,3]").Parse(),
				Tree("[3,1,null,null,2]").Parse(),
				Tree("[3,2,null,1]").Parse(),
			},
			},
		},
	})
}
