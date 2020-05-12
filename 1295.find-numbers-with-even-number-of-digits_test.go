package leetcode

import "testing"

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			Str2Slice("[12,345,2,6,7896]"),
			2,
		},
		{
			Str2Slice("[555,901,482,1771]"),
			1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.nums), func(t *testing.T) {
			res := findNumbers(test.nums)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
