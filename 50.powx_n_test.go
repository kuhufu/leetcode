package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	Run(t, myPow, []Test{
		{
			Args{
				float64(2),
				5,
			},
			Want{float64(32)},
		},
	})
}
