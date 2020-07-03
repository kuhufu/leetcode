package leetcode

/*
* 680. 验证回文字符串 Ⅱ
* https://leetcode-cn.com/problems/valid-palindrome-ii/
 */

func validPalindrome(s string) bool {
	isPalindrome := func(s string) bool {
		l := 0
		r := len(s) - 1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			break
		}
		l++
		r--
	}

	if l >= r {
		return true
	}

	return isPalindrome(s[l:r]) || isPalindrome(s[l+1:r+1])
}
