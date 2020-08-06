package leetcode

import (
	"testing"
)

func Test_canJump(t *testing.T) {

	Run(t, canJump, []Test{
		{
			Args{Slice("[5,9,3,2,1,0,2,3,3,1,0,0]").Ints()},
			Want{true},
		},
		{
			Args{Slice("[0,2,3]").Ints()},
			Want{false},
		},
		{
			Args{Slice("[0]").Ints()},
			Want{true},
		},
		{
			Args{Slice("[2,3,1,1,4]").Ints()},
			Want{true},
		},
		{
			Args{Slice("[3,2,1,0,4]").Ints()},
			Want{false},
		},
	})
}
