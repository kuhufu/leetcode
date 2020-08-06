package leetcode

import "testing"

func Test_singleNumber(t *testing.T) {
	Run(t, singleNumber, []Test{
		{
			Args{Slice("[2,2,1]").Ints()},
			Want{1},
		},
		{
			Args{Slice("[4,1,2,1,2]").Ints()},
			Want{4},
		},
	})
}
