package leetcode

import "testing"

func Test_insertIntoBST(t *testing.T) {
	Run(t, insertIntoBST, []Test{
		{
			Args{Str2Tree("[4,2,7,1,3]"), 5},
			Want{Str2Tree("[4,2,7,1,3,5]")},
		},
	})
}
