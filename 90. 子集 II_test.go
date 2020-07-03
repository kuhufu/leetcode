package leetcode

import (
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			Str2Slice("[1,2,2]"),
			nil, //返回值比较复杂，就不写了
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.nums), func(t *testing.T) {
			res := subsetsWithDup(test.nums)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
