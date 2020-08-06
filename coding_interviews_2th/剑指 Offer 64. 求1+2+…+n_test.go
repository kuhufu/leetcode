package coding_interviews_2th

import (
	"testing"
)

func Test_sumNums(t *testing.T) {
	Run(t, sumNums, []Test{
		{
			Args{10},
			Want{55},
		},
	})
}
