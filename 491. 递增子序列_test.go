package leetcode

import "testing"

func Test_findSubsequences(t *testing.T) {
	Run(t, findSubsequences, []Test{
		{
			Args{Slice("[4, 6, 7, 7]").Ints()},
			Want{Slice("[[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]").DInts()},
		},
	})
}
