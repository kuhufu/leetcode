package leetcode

import "testing"

func Test_validPalindrome(t *testing.T) {
	Run(t, validPalindrome, []Test{
		{
			Args{"atbbga"},
			Want{false},
		},
		{
			Args{"bddb"},
			Want{true},
		},
		{
			Args{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"},
			Want{true},
		},
		{
			Args{"eeccccbebaeeabebccceea"},
			Want{false},
		},
		{
			Args{"aabca"},
			Want{false},
		},
		{
			Args{"aba"},
			Want{true},
		},
		{
			Args{"abca"},
			Want{true},
		},
	})
}
