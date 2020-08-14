package leetcode

/*
* 20. 有效的括号
* https://leetcode-cn.com/problems/valid-parentheses/
 */

func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)

	for _, c := range s {
		switch c {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		case '}', ']', ')':
			n := len(stack)

			if n == 0 {
				return false
			}

			top := stack[n-1]
			stack = stack[:n-1]

			if top != byte(c) {
				return false
			}
		}
	}

	return len(stack) == 0
}
