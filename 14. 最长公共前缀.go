package leetcode

import "strings"

/*
* 14. 最长公共前缀
* https://leetcode-cn.com/problems/longest-common-prefix/
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	ans := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], ans) && len(ans)-1 >= 0 {
			ans = ans[:len(ans)-1]
		}
	}

	return ans
}
