package coding_interviews_2th

import (
	"github.com/kuhufu/leetcode"
	"sort"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			leetcode.Str2Slice("[4,1,4,6]"),
			leetcode.Str2Slice("[6,1]"),
		},
		{
			leetcode.Str2Slice("[1,2,10,4,1,4,3,3]"),
			leetcode.Str2Slice("[10,2]"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(leetcode.Str(test.nums), func(t *testing.T) {
			res := singleNumbers(test.nums)

			sort.Ints(res)
			sort.Ints(test.want)

			if !leetcode.Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
