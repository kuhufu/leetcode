package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	Run(t, canFinish, []Test{
		{
			Args{8, Str22DSlice("[[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]")},
			Want{true},
		},
		{
			Args{3, Str22DSlice("[[1,0],[2,1]]")},
			Want{true},
		},
		{
			Args{2, Str22DSlice("[[1,0]]")},
			Want{true},
		},
		{
			Args{2, Str22DSlice("[[1,0],[0,1]]")},
			Want{false},
		},
	})
}
