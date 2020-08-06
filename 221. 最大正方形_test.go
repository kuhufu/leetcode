package leetcode

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	Run(t, maximalSquare, []Test{
		{
			Args{[][]byte{
				{'0'},
				{'0'},
			}},
			Want{0},
		},
		{
			Args{[][]byte{
				{'0'},
				{'1'},
			}},
			Want{1},
		},

		{
			Args{[][]byte{
				{'0', '1'},
			}},
			Want{1},
		},
		{
			Args{[][]byte{
				{'0', '0', '0', '0', '0'},
				{'1', '0', '0', '1', '1'},
				{'1', '1', '1', '1', '0'},
			}},
			Want{1},
		},
		{
			Args{[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			Want{4},
		},
	})
}
