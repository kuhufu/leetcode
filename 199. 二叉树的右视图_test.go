package leetcode

import (
	"testing"
)

func Test_rightSideView(t *testing.T) {
	Run(t, rightSideView, []Test{
		{
			Args{Tree("[1,2,3,null,5,null,4]")},
			Want{Slice("[1, 3, 4]").Ints()},
		},
		{
			Args{Tree("[1,2,3,4]")},
			Want{Slice("[1, 3, 4]").Ints()},
		},
	})
}
