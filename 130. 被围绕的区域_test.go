package leetcode

import "testing"

func Test_solve(t *testing.T) {
	Run(t, solve, []Test{
		{
			Args{[][]byte(nil)},
			Want{[][]byte(nil)},
		},
		{
			Args{[][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			}},
			Want{[][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			}},
		},
		{
			Args{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
			Want{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
		},
	})
}
