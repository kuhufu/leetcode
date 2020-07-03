package leetcode

import (
	"fmt"
	"testing"
)

func Test_mSqrt(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			4,
			2,
		},

		{
			5,
			2,
		},
		{
			10,
			3,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.n), func(t *testing.T) {
			res := mySqrt(test.n)
			if !Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
