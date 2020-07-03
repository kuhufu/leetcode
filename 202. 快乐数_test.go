package leetcode

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := []struct {
		N    int
		want bool
	}{
		{19, true},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.N), func(t *testing.T) {
			res := isHappy2(test.N)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
