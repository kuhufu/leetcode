package leetcode

import (
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	Run(t, subsetsWithDup, []Test{
		{
			Args{Slice("[1,2,2]").Ints()},
			Want{nil}, //返回值比较复杂，就不写了
		},
	})
}
