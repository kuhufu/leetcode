package coding_interviews_2th

import (
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	Run(t, singleNumbers, []Test{
		{
			Args{Slice("[4,1,4,6]").Ints()},
			Want{Slice("[6,1]").Ints()},
		},
		{
			Args{Slice("[1,2,10,4,1,4,3,3]").Ints()},
			Want{Slice("[2,10]").Ints()},
		},
	})
}
