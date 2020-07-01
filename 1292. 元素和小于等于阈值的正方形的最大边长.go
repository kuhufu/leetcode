package leetcode

/*
* 1292. 元素和小于等于阈值的正方形的最大边长
* https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
* 二维前缀和 https://blog.csdn.net/qq_34990731/article/details/82807870
 */

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + mat[i-1][j-1]
		}
	}

	for side := min(m, n); side > 0; side-- {
		for j := side; j <= n; j++ {
			for i := side; i <= m; i++ {
				sum := dp[i][j] - dp[i][j-side] - dp[i-side][j] + dp[i-side][j-side]
				if sum <= threshold {
					return side
				}
			}
		}
	}

	return 0
}
