package leetcode

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
	Run(t, maxScoreSightseeingPair, []Test{
		{
			Args{Slice("[8,1,5,2,6]").Ints()},
			Want{11},
		},
	})
}
