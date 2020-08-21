package leetcode

import (
	"testing"
)

func Test_isSubtree(t *testing.T) {
	Run(t, isSubtree, []Test{
		{
			Args{Tree("[1,2,3]"),
				Tree("[1,2]")},
			Want{false},
		},
		{
			Args{Tree("[4,-9,5,null,-1,null,8,-6,0,7,null,null,-2,null,null,null,null,-3]"),
				Tree("[5]")},
			Want{false},
		},
		{
			Args{Tree("[3,4,5,1,2,null,null,0]"),
				Tree("[4,1,2]")},
			Want{false},
		},
		{
			Args{Tree("[-1,-4,8,-6,-2,3,9,null,-5,null,null,0,7]"),
				Tree("[3,0,5848]")},
			Want{false},
		},
		{
			Args{Tree("[3,4,5,1,2]"),
				Tree("[4,1,2]")},
			Want{true},
		},
		{
			Args{Tree("[3,4,5,1,null,2]"),
				Tree("[3,1,2]")},
			Want{false},
		},
		{
			Args{Tree("[3,4,5,1,2,null,null,null,null,0]"),
				Tree("[4,1,2]")},
			Want{false},
		},
	})
}
