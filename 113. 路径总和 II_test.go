package leetcode

import "testing"

func Test_pathSum(t *testing.T) {
	Run(t, pathSum, []Test{
		{
			Args{Tree("[1,2]"), 1},
			Want{[][]int(nil)},
		},
		{
			Args{Tree("[-2,null,-3]"), -5},
			Want{Slice("[[-2,-3]]").DInts()},
		},

		{
			Args{Tree("[5,4,8,11,null,13,4,7,2,null,null,5,1]"), 22},
			Want{Slice("[[5,4,11,2],[5,8,4,5]]").DInts()},
		},
	})
}
