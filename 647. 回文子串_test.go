package leetcode

import "testing"

func Test_countSubstrings(t *testing.T) {
	Run(t, countSubstrings, []Test{
		{
			Args{"aaaa"},
			Want{10},
		},
		{
			Args{"aaa"},
			Want{6},
		},

		{
			Args{"leetcode"},
			Want{9},
		},

		{
			Args{"abc"},
			Want{3},
		},
		{
			Args{"krggyuxvmoobsdchrgdwlopeykbdjzlln"},
			Want{36},
		},
	})
}
