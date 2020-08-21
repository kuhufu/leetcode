package leetcode

import "testing"

func Test_isBalanced(t *testing.T) {
	Run(t, isBalanced, []Test{
		{
			Args{Tree("[3,9,20,null,null,15,7]")},
			Want{true},
		},
		{
			Args{Tree("[1,2,2,3,3,null,null,4,4]")},
			Want{false},
		},
	})
}
