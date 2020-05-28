package leetcode

import "testing"

func Test_decodeString(t *testing.T) {
	Run(t, decodeString, []Test{
		{
			Args{"3[a]2[bc]"},
			Want{"aaabcbc"},
		},
		{
			Args{"3[a2[c]]"},
			Want{"accaccacc"},
		},
		{
			Args{"2[abc]3[cd]ef"},
			Want{"abcabccdcdcdef"},
		},
	})
}
