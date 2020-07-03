package leetcode

import (
	"testing"
)

func Test_numberOfSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			Str2Slice("[2,2,2,1,2,2,1,2,2,2]"),
			2,
			16,
		},
		{
			Str2Slice("[1,1,2,1,1]"),
			1,
			6,
		},
		{
			Str2Slice("[1,1,2,1,1]"),
			3,
			2,
		},
		{
			Str2Slice("[1,1,2,1,1]"),
			2,
			5,
		},
		{
			Str2Slice("[2,4,6]"),
			1,
			0,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.nums), func(t *testing.T) {
			res := numberOfSubarrays(test.nums, test.k)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
