package leetcode

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	Run(t, numIslands, []Test{
		{
			Args{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			Want{1},
		},
		{
			Args{[][]byte{
				{'1', '1', '1', '0', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '1'},
			}},
			Want{3},
		},
	})
}
