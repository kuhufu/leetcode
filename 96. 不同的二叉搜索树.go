package leetcode

/*
* 96. 不同的二叉搜索树
* https://leetcode-cn.com/problems/unique-binary-search-trees/
 */

func numTrees(n int) int {

	mem := make([]int, n+1) //记忆化
	mem[0] = 1

	var f func(int) int
	f = func(n int) int {
		if mem[n] != 0 {
			return mem[n]
		}

		sum := 0

		for i := 1; i <= n; i++ {
			sum += f(i-1) * f(n-i)
		}

		mem[n] = sum
		return sum
	}

	return f(n)
}
