package leetcode

/*
* 120. 三角形最小路径和
* *https://leetcode-cn.com/problems/triangle/
 */

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := append(triangle[n-1][:0:0], triangle[n-1]...)

	for i := n - 2; i >= 0; i-- { //从倒数第二层开始，自底向上
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}
