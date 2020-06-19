package leetcode

/*
* 125. 验证回文串
* https://leetcode-cn.com/problems/valid-palindrome/
 */

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j && !(IsAlphabet(s[i]) || IsDigital(s[i])) {
			i++
		}

		for i < j && !(IsAlphabet(s[j]) || IsDigital(s[j])) {
			j--
		}

		if i >= j {
			break
		}

		if s[i] != s[j] {
			if IsDigital(s[i]) {
				return false
			}
			if s[i]-s[j] != 32 && s[j]-s[i] != 32 {
				return false
			}
		}
		i++
		j--
	}
	return true
}
