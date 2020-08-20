package leetcode

import "testing"

func Test_updateBoard(t *testing.T) {
	Run(t, updateBoard, []Test{
		{
			Args{
				[][]byte{
					{'B', '1', 'E', '1', 'B'},
					{'B', '1', 'M', '1', 'B'},
					{'B', '1', '1', '1', 'B'},
					{'B', 'B', 'B', 'B', 'B'},
				},
				[]int{1, 2},
			},
			Want{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},
		{
			Args{
				[][]byte{
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'M', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'},
				},
				[]int{3, 0},
			},
			Want{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},
	})
}
