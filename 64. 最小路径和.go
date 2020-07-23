package leetcode

/*
* 64. 最小路径和
* https://leetcode-cn.com/problems/minimum-path-sum/
 */

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	sum := 0
	for i := 0; i < m; i++ {
		sum += grid[i][0]
		grid[i][0] = sum
	}

	sum = 0
	for i := 0; i < n; i++ {
		sum += grid[0][i]
		grid[0][i] = sum
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
