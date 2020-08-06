package coding_interviews_2th

import (
	"testing"
)

func Test_reversePairs(t *testing.T) {
	Run(t, reversePairs, []Test{
		{
			Args{Slice("[1,3,2,3,1]").Ints()},
			Want{4},
		},
		{
			Args{Slice("[7,7,5,6,7,4]").Ints()},
			Want{9},
		},
		{
			Args{Slice("[7,5,6,4]").Ints()},
			Want{5},
		},
	})
}
