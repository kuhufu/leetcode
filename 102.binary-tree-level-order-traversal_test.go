package leetcode

import "testing"

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			Str2Tree("[3,9,20,null,null,15,7]"),
			Str22DSlice(`[[3],[9,20],[15,7]]`),
		},
		{
			Str2Tree("[]"),
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.root), func(t *testing.T) {
			res := levelOrder(test.root)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
