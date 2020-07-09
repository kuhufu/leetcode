package coding_interviews_2th

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
