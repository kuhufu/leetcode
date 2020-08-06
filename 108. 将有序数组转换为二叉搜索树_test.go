package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	Run(t, sortedArrayToBST, []Test{
		{
			Args{Slice("[-10,-3,0,5,9]").Ints()},
			Want{Tree("[0,-3,9,-10,null,5]").Parse()},
		},
	})
}
