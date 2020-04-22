package leetcode

import (
	"testing"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		t    *TreeNode
		want []int
	}{
		{
			str2Tree("[1,2,3,null,5,null,4]"),
			str2Slice("[1, 3, 4]"),
		},
		{
			str2Tree("[1,2,3,4]"),
			str2Slice("[1, 3, 4]"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.t), func(t *testing.T) {
			res := rightSideView(test.t)
			if !Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
