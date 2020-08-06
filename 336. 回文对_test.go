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
			Want{Slice("[[0,1],[1,0],[2,4],[3,2]]").DInts()},
		},
	})
}

func Test_findWord(t *testing.T) {
	tree := &Node{}
	tree.make([]string{
		"a",
		"abc",
		"abcd",
		"b",
		"bc",
		"bcd",
	})

	t.Log(tree.findWord("abc"))
}
