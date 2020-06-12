package leetcode

import "testing"

func Test_pathSum(t *testing.T) {
	Run(t, pathSum, []Test{
		{
			Args{Str2Tree("[1,2]"), 1},
			Want{[][]int(nil)},
		},
		{
			Args{Str2Tree("[-2,null,-3]"), -5},
			Want{Str22DSlice("[[-2,-3]]")},
		},

		{
			Args{Str2Tree("[5,4,8,11,null,13,4,7,2,null,null,5,1]"), 22},
			Want{Str22DSlice("[[5,4,11,2],[5,8,4,5]]")},
		},
	})
}
