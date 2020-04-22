package leetcode

import (
	"fmt"
	"testing"
)

func Test_numberOfSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			str2Slice("[2,2,2,1,2,2,1,2,2,2]"),
			2,
			16,
		},
		{
			str2Slice("[1,1,2,1,1]"),
			1,
			6,
		},
		{
			str2Slice("[1,1,2,1,1]"),
			3,
			2,
		},
		{
			str2Slice("[1,1,2,1,1]"),
			2,
			5,
		},
		{
			str2Slice("[2,4,6]"),
			1,
			0,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.nums), func(t *testing.T) {
			res := numberOfSubarrays(test.nums, test.k)
			if res != test.want {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
