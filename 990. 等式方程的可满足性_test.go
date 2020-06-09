package leetcode

import "testing"

func Test_equationsPossible(t *testing.T) {
	Run(t, equationsPossible, []Test{
		{
			Args{[]string{"i!=c", "i!=f", "k==j", "g==e", "h!=e", "h==d", "j==e", "k==a", "i==h"}},
			Want{true},
		},
		{
			Args{[]string{"b==d", "c==a", "h==a", "d==d", "a==b", "h!=k", "i==h"}},
			Want{true},
		},
		{
			Args{[]string{"c==c", "f!=a", "f==b", "b==c"}},
			Want{true},
		},
		{
			Args{[]string{"a==b", "b==c", "a==c"}},
			Want{true},
		},
		{
			Args{[]string{"a==b", "b!=a"}},
			Want{false},
		},
		{
			Args{[]string{"b==a", "a==b"}},
			Want{true},
		},
		{
			Args{[]string{"a==b", "b!=c", "c==a"}},
			Want{false},
		},
	})
}
