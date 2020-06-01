package leetcode

import "testing"

func Test_kidsWithCandies(t *testing.T) {
	Run(t, kidsWithCandies, []Test{
		{
			Args{Str2Slice("[2,3,5,1,3]"), 3},
			Want{[]bool{true, true, true, false, true}},
		},
	})
}
