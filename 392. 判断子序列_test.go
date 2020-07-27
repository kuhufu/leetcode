package leetcode

import "testing"

func Test_isSubsequence(t *testing.T) {
	Run(t, isSubsequence, []Test{
		{
			Args{"abc", "ahbgdc"},
			Want{true},
		},
		{
			Args{"axc", "ahbgdc"},
			Want{false},
		},
	})
}
