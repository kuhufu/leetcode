package leetcode

import "testing"

func Test_findNumbers(t *testing.T) {
	Run(t, findNumbers, []Test{
		{
			Args{Slice("[12,345,2,6,7896]").Ints()},
			Want{2},
		},
		{
			Args{Slice("[555,901,482,1771]").Ints()},
			Want{1},
		},
	})
}
