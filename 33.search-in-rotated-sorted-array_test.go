package leetcode

import (
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			str2Slice("[4,5,6,7,0,1,2]"),
			0,
			4,
		},
		{
			str2Slice("[4,5,6,7,0,1,2]"),
			3,
			-1,
		},
	}

	for _, test := range tests {
		t.Run(Str(test.nums), func(t *testing.T) {
			res := search(test.nums, test.target)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
