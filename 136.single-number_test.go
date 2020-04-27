package leetcode

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			str2Slice("[2,2,1]"),
			1,
		},
		{
			str2Slice("[4,1,2,1,2]"),
			4,
		},
	}

	for _, test := range tests {
		t.Run(Str(test.nums), func(t *testing.T) {
			res := singleNumber(test.nums)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
