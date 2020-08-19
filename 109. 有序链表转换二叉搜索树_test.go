package leetcode

import "testing"

func Test_sortedListToBST(t *testing.T) {
	Run(t, sortedListToBST, []Test{
		{
			Args{List("[-10, -3, 0, 5, 9]").Parse()},
			Want{Tree("[0,-10,5,null,-3,null,9]").Parse()},
		},
	})
}
