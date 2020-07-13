package leetcode

import "testing"

func Test_intersect(t *testing.T) {
	Run(t, intersect, []Test{
		{
			Args{Str2Slice("[1,2,2,1]"), Str2Slice("[2,2]")},
			Want{Str2Slice("[2,2]")},
		},
		{
			Args{Str2Slice("[4,9,5]"), Str2Slice("[9,4,9,8,4]")},
			Want{Str2Slice("[9,4]")},
		},
	})
}
