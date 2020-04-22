package leetcode

/*
* 200. 岛屿数量
* https://leetcode-cn.com/problems/number-of-islands/submissions/
 */

const (
	NoVisitedIsland = '1'
	VisitedIsland   = '2'
)

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == NoVisitedIsland {
				ans++
				grid[i][j] = VisitedIsland

				dfs1(i-1, j, grid)
				dfs1(i+1, j, grid)
				dfs1(i, j-1, grid)
				dfs1(i, j+1, grid)
			}
		}
	}

	return ans
}

func dfs1(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == NoVisitedIsland {
		grid[i][j] = VisitedIsland

		dfs1(i-1, j, grid)
		dfs1(i+1, j, grid)
		dfs1(i, j-1, grid)
		dfs1(i, j+1, grid)
	}
}
