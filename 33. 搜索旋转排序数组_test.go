package leetcode

import (
	"testing"
)

func Test_search(t *testing.T) {
	Run(t, search, []Test{
		{
			Args{
				Slice("[4,5,6,7,0,1,2]").Ints(),
				0,
			},
			Want{4},
		},
		{
			Args{
				Slice("[4,5,6,7,0,1,2]").Ints(),
				3,
			},
			Want{-1},
		},
	})
}
