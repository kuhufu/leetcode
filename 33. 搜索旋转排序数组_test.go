package leetcode

import (
	"testing"
)

func Test_search(t *testing.T) {
	Run(t, search, []Test{
		{
			Args{
				Str2Slice("[4,5,6,7,0,1,2]"),
				0,
			},
			Want{4},
		},
		{
			Args{
				Str2Slice("[4,5,6,7,0,1,2]"),
				3,
			},
			Want{-1},
		},
	})
}
