package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	Run(t, canFinish, []Test{
		{
			Args{8, Slice("[[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]").DInts()},
			Want{true},
		},
		{
			Args{3, Slice("[[1,0],[2,1]]").DInts()},
			Want{true},
		},
		{
			Args{2, Slice("[[1,0]]").DInts()},
			Want{true},
		},
		{
			Args{2, Slice("[[1,0],[0,1]]").DInts()},
			Want{false},
		},
	})
}
