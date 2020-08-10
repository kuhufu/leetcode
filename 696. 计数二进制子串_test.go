package leetcode

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
	Run(t, countBinarySubstrings, []Test{
		{
			Args{"00110011"},
			Want{6},
		},
		{
			Args{"10101"},
			Want{4},
		},
	})
}
