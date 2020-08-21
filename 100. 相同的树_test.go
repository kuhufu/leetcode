package leetcode

import "testing"

func Test_isSameTree(t *testing.T) {
	Run(t, isSameTree, []Test{
		{
			Args{
				Tree("[1,2,3]"),
				Tree("[1,2,3]")},
			Want{true},
		},
		{
			Args{
				Tree("[1,2]"),
				Tree("[1,null,2]")},
			Want{false},
		},
	})
}
