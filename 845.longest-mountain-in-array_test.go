package leetcode

import "testing"

func Test_longestMountain(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			Str2Slice("[2,1,4,7,3,2,5,2,1]"),
			5,
		},
		{
			Str2Slice("[2,1,4,7,3,2,5]"),
			5,
		},
		{
			Str2Slice("[2,2,2]"),
			0,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.A), func(t *testing.T) {
			res := longestMountain(test.A)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
