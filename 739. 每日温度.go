package leetcode

/*
* 739. 每日温度
* https://leetcode-cn.com/problems/daily-temperatures/
 */

/*
* tag: 最小栈
 */

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := make([]int, 0, len(T))

	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] { //当前气温高于栈顶气温
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1] //出栈
		}

		stack = append(stack, i)
	}

	return ans
}
