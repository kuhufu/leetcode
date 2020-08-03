package leetcode

import "testing"

func Test_addStrings(t *testing.T) {
	Run(t, addStrings, []Test{
		{
			Args{"0", "0"},
			Want{"0"},
		},
		{
			Args{"1", "9"},
			Want{"10"},
		},
		{
			Args{"9", "99"},
			Want{"108"},
		},
	})
}
