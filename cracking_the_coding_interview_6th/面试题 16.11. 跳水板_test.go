package cracking_the_coding_interview_6th

import (
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
			Want{Slice("[3,4,5,6]").Ints()},
		},
	})
}
