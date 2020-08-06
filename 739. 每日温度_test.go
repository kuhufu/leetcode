package leetcode

import "testing"

func Test_dailyTemperatures(t *testing.T) {
	Run(t, dailyTemperatures, []Test{
		{
			Args{Slice("[73, 74, 75, 71, 69, 72, 76, 73]").Ints()},
			Want{Slice("[1, 1, 4, 2, 1, 1, 0, 0]").Ints()},
		},
	})
}
