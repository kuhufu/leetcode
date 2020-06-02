package cracking_the_coding_interview_6th

import (
	"github.com/kuhufu/leetcode"
	"testing"
)

func Test_sumNums(t *testing.T) {
	leetcode.Run(t, sumNums, []leetcode.Test{
		{
			leetcode.Args{10},
			leetcode.Want{55},
		},
	})
}
