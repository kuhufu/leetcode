package leetcode

import "testing"

func Test_robIII(t *testing.T) {
	Run(t, robIII, []Test{
		{
			Args{Tree("[4,1,null,2,null,3]")},
			Want{7},
		},
		{
			Args{Tree("[3,2,3,null,3,null,1]")},
			Want{7},
		},
		{
			Args{Tree("[3,4,5,1,3,null,1]")},
			Want{9},
		},
	})
}
