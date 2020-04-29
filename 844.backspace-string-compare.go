package leetcode

/*
* 844. 比较含退格的字符串
* https://leetcode-cn.com/problems/backspace-string-compare/submissions/
 */

/*
* 难度: 简单
 */

//栈
func backspaceCompare(S string, T string) bool {
	f := func(S string) string {
		s := make([]byte, 0, len(S))
		for _, v := range []byte(S) {
			if v == '#' {
				if len(s) != 0 {
					s = s[:len(s)-1]
				}
			} else {
				s = append(s, v)
			}
		}
		return string(s)
	}

	return f(S) == f(T)
}
