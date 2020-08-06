package leetcode

import "testing"

func Test_kidsWithCandies(t *testing.T) {
	Run(t, kidsWithCandies, []Test{
		{
			Args{Slice("[2,3,5,1,3]").Ints(), 3},
			Want{[]bool{true, true, true, false, true}},
		},
	})
}
