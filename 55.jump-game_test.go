package leetcode

import (
	"fmt"
	"testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, true},
		{[]int{0, 2, 3}, false},
		{[]int{0}, true},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.nums), func(t *testing.T) {
			out := canJump(test.nums)
			if test.want != out {
				t.Errorf("want %v, but got %v", test.want, out)
			}
		})
	}
}
