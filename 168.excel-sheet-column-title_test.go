package leetcode

import (
	"fmt"
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
		t.Run(fmt.Sprint(test.N), func(t *testing.T) {
			out := convertToTitle(test.N)
			if test.want != out {
				t.Errorf("want %v, but got %v", test.want, out)
			}
		})
	}
}
