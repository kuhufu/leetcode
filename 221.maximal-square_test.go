package leetcode

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			[][]byte{
				{'0'},
				{'0'},
			},
			0,
		},
		{
			[][]byte{
				{'0'},
				{'1'},
			},
			1,
		},

		{
			[][]byte{
				{'0', '1'},
			},
			1,
		},
		{
			[][]byte{
				{'0', '0', '0', '0', '0'},
				{'1', '0', '0', '1', '1'},
				{'1', '1', '1', '1', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.matrix), func(t *testing.T) {
			res := maximalSquare(test.matrix)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
