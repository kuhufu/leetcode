package leetcode

import "testing"

func Test_kthSmallest(t *testing.T) {
	Run(t, kthSmallest, []Test{
		{
			Args{Slice("[[1,5,9],[10,11,13],[12,13,15]]").DInts(), 8},
			Want{13},
		},
	})
}
