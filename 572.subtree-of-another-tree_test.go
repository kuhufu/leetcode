package leetcode

import "testing"

func Test_isSubtree(t *testing.T) {
	tests := []struct {
		s    *TreeNode
		t    *TreeNode
		want bool
	}{
		{
			Str2Tree("[1,2,3]"),
			Str2Tree("[1,2]"),
			false,
		},
		{
			Str2Tree("[4,-9,5,null,-1,null,8,-6,0,7,null,null,-2,null,null,null,null,-3]"),
			Str2Tree("[5]"),
			false,
		},
		{
			Str2Tree("[3,4,5,1,2,null,null,0]"),
			Str2Tree("[4,1,2]"),
			false,
		},
		{
			Str2Tree("[-1,-4,8,-6,-2,3,9,null,-5,null,null,0,7]"),
			Str2Tree("[3,0,5848]"),
			false,
		},
		{
			Str2Tree("[3,4,5,1,2]"),
			Str2Tree("[4,1,2]"),
			true,
		},
		{
			Str2Tree("[3,4,5,1,null,2]"),
			Str2Tree("[3,1,2]"),
			false,
		},
		{
			Str2Tree("[3,4,5,1,2,null,null,null,null,0]"),
			Str2Tree("[4,1,2]"),
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.s, test.t), func(t *testing.T) {
			res := isSubtree2(test.s, test.t)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
