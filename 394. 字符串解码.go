package leetcode

/*
* 394. 字符串解码
* https://leetcode-cn.com/problems/decode-string/
 */

/*
* 难度： 中等
 */
func decodeString(s string) string {
	string2, _ := decodeString2(s)
	return string2
}

func decodeString2(s string) (string, int) {
	if len(s) == 0 {
		return "", 0
	}

	count := 0
	ans := ""
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		switch {
		case IsDigital(c):
			count = count*10 + int(c-'0')
		case IsAlphabet(c):
			ans += s[i : i+1]
		case c == '[':
			tmp, idx := decodeString2(s[i+1:])
			for i := 0; i < count; i++ {
				ans += tmp
			}
			i += idx + 1
			count = 0
		case c == ']':
			return ans, i
		}
	}

	return ans, i
}
