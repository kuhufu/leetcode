package leetcode

import (
	"testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			Str2Slice("[5,9,3,2,1,0,2,3,3,1,0,0]"),
			true,
		},
		{
			Str2Slice("[0,2,3]"),
			false,
		},
		{
			Str2Slice("[0]"),
			true,
		},
		{
			Str2Slice("[2,3,1,1,4]"),
			true,
		},
		{
			Str2Slice("[3,2,1,0,4]"),
			false,
		},
	}

	for _, test := range tests {
		t.Run(Str(test.nums), func(t *testing.T) {
			res := canJump(test.nums)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
