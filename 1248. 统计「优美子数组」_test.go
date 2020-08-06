package leetcode

import (
	"testing"
)

func Test_numberOfSubarrays(t *testing.T) {
	Run(t, numberOfSubarrays, []Test{
		{
			Args{Slice("[2,2,2,1,2,2,1,2,2,2]").Ints(),
				2},
			Want{16},
		},
		{
			Args{Slice("[1,1,2,1,1]").Ints(),
				1},
			Want{6},
		},
		{
			Args{Slice("[1,1,2,1,1]").Ints(),
				3},
			Want{2},
		},
		{
			Args{Slice("[1,1,2,1,1]").Ints(),
				2},
			Want{5},
		},
		{
			Args{Slice("[2,4,6]").Ints(),
				1},
			Want{0},
		},
	})
}
