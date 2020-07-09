package leetcode

import "math"

/*
* 1392. 最长快乐前缀
* https://leetcode-cn.com/problems/longest-happy-prefix/
 */

/*
* 难度: 困难
* 字符串哈希 可以用来快速判断 字符串是否含有子串
 */

func longestPrefix(s string) string {
	return longestPrefix1(s)
}

//方法一：Rabin-Karp 字符串编码，还可以理解
func longestPrefix1(s string) string {
	n := len(s)
	prefix, suffix := 0, 0
	base := 31
	mod := math.MaxInt32
	mul := 1

	happy := 0

	for i := 1; i < n; i++ {
		prefix = (prefix*base + int(s[i-1]-'a')) % mod
		suffix = (suffix + int(s[n-i]-'a')*mul) % mod
		if prefix == suffix {
			happy = i
		}
		mul = mul * base % mod
	}

	return s[:happy]
}

//方法二：KMP 算法，难以理解
func longestPrefix2(s string) string {
	n := len(s)

	fail := make([]int, n)
	for i := 0; i < len(fail); i++ {
		fail[i] = -1
	}

	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && s[j+1] != s[i] {
			j = fail[j]
		}
		if s[j+1] == s[i] {
			fail[i] = j + 1
		}
	}

	return s[:fail[n-1]+1]
}
