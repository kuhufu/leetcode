package cracking_the_coding_interview_6th

import (
	"github.com/kuhufu/leetcode"
	"testing"
)

func Test_waysToChange(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 4},
		{1, 1},
		{5, 2},
	}

	for _, test := range tests {
		test := test
		t.Run(leetcode.Str(test.n), func(t *testing.T) {
			res := waysToChange(test.n)
			if !leetcode.Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
