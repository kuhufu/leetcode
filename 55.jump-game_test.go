package leetcode

import (
	"testing"
)

func Test_canJump(t *testing.T) {

	Run(t, canJump, []Test{
		{
			Args{Str2Slice("[5,9,3,2,1,0,2,3,3,1,0,0]")},
			Want{true},
		},
		{
			Args{Str2Slice("[0,2,3]")},
			Want{false},
		},
		{
			Args{Str2Slice("[0]")},
			Want{true},
		},
		{
			Args{Str2Slice("[2,3,1,1,4]")},
			Want{true},
		},
		{
			Args{Str2Slice("[3,2,1,0,4]")},
			Want{false},
		},
	})
}
