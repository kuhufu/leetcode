package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	Run(t, minSubArrayLen, []Test{
		{
			Args{3, Slice("[1,1]").Ints()},
			Want{0},
		},
		{
			Args{15, Slice("[5,1,3,5,10,7,4,9,2,8]").Ints()},
			Want{2},
		},
		{
			Args{7, Slice("[2,3,1,2,4,3]").Ints()},
			Want{2},
		},
	})
}
