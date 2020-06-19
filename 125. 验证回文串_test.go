package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	Run(t, isPalindrome, []Test{
		{
			Args{"A man, a plan, a canal: Panama"},
			Want{true},
		},
		{
			Args{"race a car"},
			Want{false},
		},
		{
			Args{"0P"},
			Want{false},
		},
	})
}
