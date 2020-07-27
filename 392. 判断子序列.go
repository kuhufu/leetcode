package leetcode

/*
* 392. 判断子序列
* https://leetcode-cn.com/problems/is-subsequence/
 */

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i != len(s) && j != len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}
