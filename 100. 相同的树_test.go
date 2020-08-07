package leetcode

import "testing"

func Test_isSameTree(t *testing.T) {
	Run(t, isSameTree, []Test{
		{
			Args{
				Tree("[1,2,3]").Parse(),
				Tree("[1,2,3]").Parse()},
			Want{true},
		},
		{
			Args{
				Tree("[1,2]").Parse(),
				Tree("[1,null,2]").Parse()},
			Want{false},
		},
	})
}
