package leetcode

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	Run(t, levelOrder, []Test{
		{
			Args{Tree("[3,9,20,null,null,15,7]").Parse()},
			Want{Slice(`[[3],[9,20],[15,7]]`).DInts()},
		},
		{
			Args{Tree("[]").Parse()},
			Want{nil},
		},
	})
}
