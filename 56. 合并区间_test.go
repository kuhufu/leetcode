package leetcode

import (
	"testing"
)

func Test_merge(t *testing.T) {
	Run(t, merge, []Test{
		{
			Args{Slice("[[1,3],[2,6],[8,10],[15,18]]").DInts()},
			Want{Slice("[[1,6],[8,10],[15,18]]").DInts()},
		},
		{
			Args{Slice("[[1,4],[4,5]]").DInts()},
			Want{Slice("[[1,5]]").DInts()},
		},
		{
			Args{Slice("[[2,3],[1,4]]").DInts()},
			Want{Slice("[[1,4]]").DInts()},
		},
	})
}
