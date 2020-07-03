package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	Run(t, sortedArrayToBST, []Test{
		{
			Args{Str2Slice("[-10,-3,0,5,9]")},
			Want{Str2Tree("[0,-3,9,-10,null,5]")},
		},
	})
}
