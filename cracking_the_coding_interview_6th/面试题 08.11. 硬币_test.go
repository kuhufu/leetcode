package cracking_the_coding_interview_6th

import (
	"github.com/kuhufu/leetcode"
	"testing"
)

func Test_waysToChange(t *testing.T) {
	leetcode.Run(t, waysToChange, []Test{
		{
			Args{10},
			Want{4},
		},
		{
			Args{1},
			Want{1},
		},
		{
			Args{5},
			Want{2},
		},
	})
}
