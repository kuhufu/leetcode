package leetcode

import "testing"

func Test_suggestedProducts(t *testing.T) {
	Run(t, suggestedProducts, []Test{
		{
			Args{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"},
			Want{[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			}},
		},
		{
			Args{[]string{"havana"}, "havana"},
			Want{[][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			}},
		},
	})
}
