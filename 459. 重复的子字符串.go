package leetcode

import "strings"

/*
* 459. 重复的子字符串
* https://leetcode-cn.com/problems/repeated-substring-pattern/
 */

func repeatedSubstringPattern(s string) bool {
	n := len(s)

	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}

		prefix := s[:i]
		tmp := s
		for i <= len(tmp) && strings.HasPrefix(tmp, prefix) {
			tmp = tmp[i:]
		}

		if len(tmp) == 0 {
			return true
		}
	}

	return false
}
