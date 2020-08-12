package leetcode

import "testing"

func Test_cloneGraph(t *testing.T) {
	Run(t, cloneGraph, []Test{
		{
			Args{Graph("[[2,4],[1,3],[2,4],[1,3]]").Parse()},
			Want{Graph("[[2,4],[1,3],[2,4],[1,3]]").Parse()},
		},
	})
}
