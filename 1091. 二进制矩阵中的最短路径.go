package leetcode

/*
* 1091. 二进制矩阵中的最短路径
* https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
* 难度: 中等
 */

/*
* 提交通过: true
 */

func shortestPathBinaryMatrix(grid [][]int) int {
	dir := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	if grid[0][0] != 0 {
		return -1
	}

	N := len(grid)

	grid[0][0] = 2
	queue := [][]int{{0, 0}}

	ans := 1
	for len(queue) != 0 {
		size := len(queue)
		for _, idx := range queue[:size] {
			i := idx[0]
			j := idx[1]

			if i == N-1 && j == N-1 {
				return ans
			}

			for _, d := range dir {
				i2 := i + d[0]
				j2 := j + d[1]
				if i2 < 0 || i2 >= N || j2 < 0 || j2 >= N {
					continue
				}

				if grid[i2][j2] != 0 {
					continue
				}

				grid[i2][j2] = 2
				queue = append(queue, []int{i2, j2})
			}
		}
		queue = queue[size:]
		ans++
	}

	return -1
}
