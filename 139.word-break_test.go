package leetcode

import "testing"

func Test2_wordBreak(t *testing.T) {
	Run(t, wordBreak, []Test{
		{
			Args{
				"leetcode",
				[]string{"leet", "code"},
			},
			Want{true},
		},
		{
			Args{
				"applepenapple",
				[]string{"apple", "pen"},
			},
			Want{true},
		},
		{
			Args{
				"catsandog",
				[]string{"cats", "dog", "sand", "and", "cat"},
			},
			Want{false},
		},
	})
}
