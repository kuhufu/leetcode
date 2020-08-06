package coding_interviews_2th

import (
	"testing"
)

func Test_translateNum(t *testing.T) {
	Run(t, translateNum, []Test{
		{
			Args{25},
			Want{2},
		},
		{
			Args{12258},
			Want{5},
		},
		{
			Args{102},
			Want{2},
		},
	})
}
