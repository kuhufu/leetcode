package leetcode

import "testing"

func Test_palindromePairs(t *testing.T) {
	Run(t, palindromePairs, []Test{
		{
			Args{Slice(`["bat", "tab", "cat"]`).Strings()},
			Want{Slice("[[0,1],[1,0]]").DInts()},
		},
		{
			Args{Slice(`["abcd", "dcba", "lls", "s", "sssll"]`).Strings()},
			Want{Slice("[[0,1],[1,0],[2,4]],[3,2]").DInts()},
		},
	})
}
