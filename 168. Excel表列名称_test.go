package leetcode

import (
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		N    int
		want string
	}{
		{1, "A"},
		{26, "Z"},
		{28, "AB"},
		{52, "AZ"},
		{701, "ZY"},
		{26 * 26, "YZ"},
		{26*26 + 1, "ZA"},
		{26 * 26 * 26, "YYZ"},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.N), func(t *testing.T) {
			res := convertToTitle(test.N)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
