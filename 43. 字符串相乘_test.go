package leetcode

import "testing"

func Test_multiply(t *testing.T) {
	Run(t, multiply, []Test{
		{
			Args{"0", "52"},
			Want{"0"},
		},
		{
			Args{"9999", "0"},
			Want{"0"},
		},
		{
			Args{"9", "9"},
			Want{"81"},
		},
		{
			Args{"123", "456"},
			Want{"56088"},
		},
	})
}

func Benchmark_multiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiply("45612312321", "45612312321")
	}
}

func Test_add(t *testing.T) {
	Run(t, add, []Test{
		{
			Args{"0", "0"},
			Want{"0"},
		},
		{
			Args{"1", "9"},
			Want{"10"},
		},
		{
			Args{"9", "99"},
			Want{"108"},
		},
		{
			Args{"16", "2"},
			Want{"18"},
		},
	})
}
