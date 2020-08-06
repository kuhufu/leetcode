package leetcode

import "testing"

func Test_recoverFromPreorder(t *testing.T) {
	Run(t, recoverFromPreorder, []Test{
		{
			Args{"1"},
			Want{Tree("[1]").Parse()},
		},
		{
			Args{"1-2--3---4-5--6---7"},
			Want{Tree("[1,2,5,3,null,6,null,4,null,7]").Parse()},
		},
		{
			Args{"1-401--349---90--88"},
			Want{Tree("[1,401,null,349,88,90]").Parse()},
		},
	})
}
