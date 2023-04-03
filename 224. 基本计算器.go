package leetcode

import "fmt"

/*
* 224. 基本计算器
* https://leetcode.cn/problems/basic-calculator/
 */

/*
* 难度: 困难
 */

func calculate(s string) int {
	var nums []int
	var ops []byte

	for i := 0; i < len(s); i++ {
		c := s[i]

		if isDigit(c) {
			var num int
			for ; i < len(s) && isDigit(s[i]); i++ {
				num = num*10 + int(c-'0')
			}
			nums = append(nums, num)
			i--
			continue
		}

		switch c {
		case '(':
		case ')':
			for {
				prevOp := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if prevOp == '(' {
					break
				}
				num := eval(nums, prevOp)
				nums = nums[:len(nums)-2]
				nums = append(nums, num)
			}
			continue
		case '+', '-', '*', '/':
			if len(ops) == 0 || ops[len(ops)-1] == '(' {
				break
			}

			prevOp := ops[len(ops)-1]
			if Priority(prevOp) >= Priority(c) {
				num := eval(nums, prevOp)

				nums = nums[:len(nums)-2]
				ops = ops[:len(ops)-1]

				nums = append(nums, num)
				i--
				continue
			}
		}
		ops = append(ops, c)
	}

	fmt.Println(nums)
	fmt.Println(string(ops))

	for len(nums) != 1 {
		op := ops[len(ops)-1]
		num := eval(nums, op)

		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]

		nums = append(nums, num)
	}

	return nums[0]
}

func eval(nums []int, op byte) int {
	x := nums[len(nums)-2]
	y := nums[len(nums)-1]

	switch op {
	case '+':
		return x + y
	case '-':
		return x - y
	case '*':
		return x * y
	case '/':
		return x / y
	}

	return 0
}

func Priority(c byte) int {
	switch c {
	case '(':
		return 0
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}

	return 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
