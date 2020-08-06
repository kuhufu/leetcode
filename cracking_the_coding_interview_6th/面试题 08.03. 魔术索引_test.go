package cracking_the_coding_interview_6th

import (
	"testing"
)

func Test_findMagicIndex(t *testing.T) {
	Run(t, findMagicIndex, []Test{
		{
			Args{Slice("[0, 0, 2]").Ints()},
			Want{0},
		},
		{
			Args{Slice("[0, 2, 3, 4, 5]").Ints()},
			Want{0},
		},
		{
			Args{Slice("[1, 1, 1]").Ints()},
			Want{1},
		},
	})
}
