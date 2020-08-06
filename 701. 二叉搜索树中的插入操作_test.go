package leetcode

import "testing"

func Test_insertIntoBST(t *testing.T) {
	Run(t, insertIntoBST, []Test{
		{
			Args{Tree("[4,2,7,1,3]").Parse(), 5},
			Want{Tree("[4,2,7,1,3,5]").Parse()},
		},
	})
}
