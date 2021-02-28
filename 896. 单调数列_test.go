package leetcode

import "testing"

func Test_isMonotonic(t *testing.T) {
	Run(t, isMonotonic, []Test{
		{
			Args{Slice("[3,1,9]").Ints()},
			Want{false},
		},
		{
			Args{Slice("[1,2,2,3]").Ints()},
			Want{true},
		},
	})
}
