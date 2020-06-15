package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	Run(t, longestCommonPrefix, []Test{
		{
			Args{[]string{"flower", "flow", "flight"}},
			Want{"fl"},
		},
		{
			Args{[]string{"dog", "racecar", "car"}},
			Want{""},
		},
	})
}
