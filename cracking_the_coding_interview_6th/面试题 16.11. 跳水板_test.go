package cracking_the_coding_interview_6th

import (
	. "github.com/kuhufu/leetcode"
	"testing"
)

func Test_divingBoard(t *testing.T) {
	Run(t, divingBoard, []Test{
		{
			Args{1, 1, 0},
			Want{[]int(nil)},
		},
		{
			Args{1, 2, 3},
			Want{Str2Slice("[3,4,5,6]")},
		},
	})
}
