package leetcode

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	Run(t, repeatedSubstringPattern, []Test{
		{
			Args{"abab"},
			Want{true},
		},
		{
			Args{"aba"},
			Want{false},
		},
	})
}
