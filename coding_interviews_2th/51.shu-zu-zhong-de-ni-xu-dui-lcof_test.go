package coding_interviews_2th

import (
	"github.com/kuhufu/leetcode"
	"testing"
)

func Test_reversePairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 3, 2, 3, 1},
			4,
		},
		{
			[]int{7, 7, 5, 6, 7, 4},
			9,
		},
		{
			[]int{7, 5, 6, 4},
			5,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(leetcode.Str(test.nums), func(t *testing.T) {
			res := reversePairs(test.nums)
			if !leetcode.Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
