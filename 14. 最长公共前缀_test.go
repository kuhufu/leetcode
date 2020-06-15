package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	Run(t, longestCommonPrefix, []Test{
		{
			Args{[]string{}},
			Want{""},
		},
		{
			Args{[]string{"abc"}},
			Want{"abc"},
		},
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
