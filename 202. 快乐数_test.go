package leetcode

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	Run(t, isHappy, []Test{
		{
			Args{19},
			Want{true},
		},
	})
}
