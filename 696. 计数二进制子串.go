package leetcode

/*
* 696. 计数二进制子串
* https://leetcode-cn.com/problems/count-binary-substrings/
 */

func countBinarySubstrings(s string) int {
	ans := 0
	for i := 1; i < len(s); {
		if s[i] == s[i-1] {
			i++
			continue
		}

		cur := i
		for j := 0; cur-j-1 >= 0 && cur+j < len(s) && s[cur] == s[cur+j] && s[cur-j-1] != s[cur+j]; j++ {
			ans++
			i++
		}
	}

	return ans
}
