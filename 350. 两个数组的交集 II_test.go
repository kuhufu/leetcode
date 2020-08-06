package leetcode

import "testing"

func Test_intersect(t *testing.T) {
	Run(t, intersect, []Test{
		{
			Args{Slice("[1,2,2,1]"), Slice("[2,2]").Ints()},
			Want{Slice("[2,2]").Ints()},
		},
		{
			Args{Slice("[4,9,5]"), Slice("[9,4,9,8,4]").Ints()},
			Want{Slice("[9,4]").Ints()},
		},
	})
}
