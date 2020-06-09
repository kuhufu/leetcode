package cracking_the_coding_interview_6th

import "strconv"

/*
* 面试题46. 把数字翻译成字符串
* https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
 */

func translateNum(num int) int {
	s := strconv.Itoa(num)
	N := len(s)
	mem := make([]int, N)
	mem[N-1] = 1

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == N {
			return 1
		}
		if mem[i] != 0 {
			return mem[i]
		}

		l := 0
		r := 0

		if i < N {
			l = dfs(i + 1)
		}

		if i+1 < N && s[i:i+2] >= "10" && s[i:i+2] <= "25" {
			r = dfs(i + 2)
		}
		mem[i] = l + r
		return mem[i]
	}

	dfs(0)

	return mem[0]
}
