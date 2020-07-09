package leetcode

import "testing"

func Test_longestPrefix(t *testing.T) {
	Run(t, longestPrefix, []Test{
		{
			Args{"level"},
			Want{"l"},
		},
		{
			Args{"ababab"},
			Want{"abab"},
		},
		{
			Args{"leetcodeleet"},
			Want{"leet"},
		},
		{
			Args{"a"},
			Want{""},
		},
	})
}
