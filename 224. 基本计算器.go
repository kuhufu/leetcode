package leetcode

/*
* 224. 基本计算器
* https://leetcode.cn/problems/basic-calculator/
 */

/*
* 难度: 困难
 */

func calculate(s string) int {
	return helper(s, 0)
}

func helper(s string, i int) int {
	var num int
	var sign = '+'
	var stack []int
	for ; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		}

		if isDigit(c) {
			num = num*10 + int(c-'0')
		}

		if !isDigit(c) || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			case '(':
				helper(s[i+1:], i)
			case ')':
				return sumInt(stack)
			}

			sign = int32(c)
			num = 0
		}
	}

	return sumInt(stack)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
