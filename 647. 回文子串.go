package leetcode

/*
* 647. 回文子串
* https://leetcode-cn.com/problems/palindromic-substrings/
 */

func countSubstrings(s string) int {
	n := len(s)

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	ans := 0

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[j] == s[i] {
				dp[i][j] = j == i || j-i == 1 || dp[i+1][j-1]
				if dp[i][j] {
					ans++
				}
			}
		}
	}

	return ans
}

func countSubstrings1(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += foo1(s, i)
	}

	return ans
}

func foo1(s string, idx int) int {
	ret := 1

	l, r := idx-1, idx+1

	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			ret++
		} else {
			break
		}

		l--
		r++
	}

	l, r = idx-1, idx+1

	if l >= 0 && s[l] == s[idx] {
		ret++
		l--

		for l >= 0 && r < len(s) {
			if s[l] == s[r] {
				ret++
			} else {
				break
			}

			l--
			r++
		}
	}

	return ret
}
