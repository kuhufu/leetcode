package leetcode

import (
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	Run(t, convertToTitle, []Test{
		{
			Args{1},
			Want{"A"},
		},
		{
			Args{26},
			Want{"Z"},
		},
		{
			Args{28},
			Want{"AB"},
		},
		{
			Args{52},
			Want{"AZ"},
		},
		{
			Args{701},
			Want{"ZY"},
		},
		{
			Args{26 * 26},
			Want{"YZ"},
		},
		{
			Args{26*26 + 1},
			Want{"ZA"},
		},
		{
			Args{26 * 26 * 26},
			Want{"YYZ"},
		},
	})
}
