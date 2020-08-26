package leetcode

import "testing"

func Test_letterCombinations(t *testing.T) {
	Run(t, letterCombinations, []Test{
		{
			Args{""},
			Want{[]string(nil)},
		},
		{
			Args{"23"},
			Want{Slice(`["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]`).Strings()},
		},
	})
}
