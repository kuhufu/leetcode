package leetcode

import "testing"

func Test_kthSmallest(t *testing.T) {
	Run(t, kthSmallest, []Test{
		{
			Args{Str22DSlice("[[1,5,9],[10,11,13],[12,13,15]]"), 8},
			Want{13},
		},
	})
}
