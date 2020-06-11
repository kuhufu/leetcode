package leetcode

import "testing"

func Test_dailyTemperatures(t *testing.T) {
	Run(t, dailyTemperatures, []Test{
		{
			Args{Str2Slice("[73, 74, 75, 71, 69, 72, 76, 73]")},
			Want{Str2Slice("[1, 1, 4, 2, 1, 1, 0, 0]")},
		},
	})
}
