package leetcode

import "testing"

func Test_recoverFromPreorder(t *testing.T) {
	Run(t, recoverFromPreorder, []Test{
		{
			Args{"1"},
			Want{Str2Tree("[1]")},
		},
		{
			Args{"1-2--3---4-5--6---7"},
			Want{Str2Tree("[1,2,5,3,null,6,null,4,null,7]")},
		},
		{
			Args{"1-401--349---90--88"},
			Want{Str2Tree("[1,401,null,349,88,90]")},
		},
	})
}
