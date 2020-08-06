package leetcode

import "testing"

func Test_isBipartite(t *testing.T) {
	Run(t, isBipartite, []Test{
		{
			Args{Slice("[[],[2,4,6],[1,4,8,9],[7,8],[1,2,8,9],[6,9],[1,5,7,8,9],[3,6,9],[2,3,4,6,9],[2,4,5,6,7,8]]").DInts()},
			Want{false},
		},
		{
			Args{Slice("[[1,3], [0,2], [1,3], [0,2]]").DInts()},
			Want{true},
		},
		{
			Args{Slice("[[1,2,3], [0,2], [0,1,3], [0,2]]").DInts()},
			Want{false},
		},
	})
}
