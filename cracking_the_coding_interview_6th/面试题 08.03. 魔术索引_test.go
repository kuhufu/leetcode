package cracking_the_coding_interview_6th

import (
	. "github.com/kuhufu/leetcode"
	"testing"
)

func Test_findMagicIndex(t *testing.T) {
	Run(t, findMagicIndex, []Test{
		{
			Args{Str2Slice("[0, 0, 2]")},
			Want{0},
		},
		{
			Args{Str2Slice("[0, 2, 3, 4, 5]")},
			Want{0},
		},
		{
			Args{Str2Slice("[1, 1, 1]")},
			Want{1},
		},
	})
}
