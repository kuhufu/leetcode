package coding_interviews_2th

import (
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	Run(t, spiralOrder, []Test{
		{
			Args{Slice("[[1,2,3,4],[5,6,7,8],[9,10,11,12]]").DInts()},
			Want{Slice("[1,2,3,4,8,12,11,10,9,5,6,7]").Ints()},
		},
		{
			Args{Slice("[[1,2,3],[4,5,6],[7,8,9]]").DInts()},
			Want{Slice("[1,2,3,6,9,8,7,4,5]").Ints()},
		},
	})
}
