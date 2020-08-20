package leetcode

/*
* 529. 扫雷游戏
* https://leetcode-cn.com/problems/minesweeper/
 */

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	i, j := click[0], click[1]
	if board[i][j] == 'M' {
		board[i][j] = 'X'
		return board
	}

	var dfs func([][]byte, int, int)
	dfs = func(board [][]byte, i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}

		if board[i][j] != 'E' {
			return
		}

		count := count(board, i, j)
		if count != 0 {
			board[i][j] = '0' + byte(count)
			return
		}

		board[i][j] = 'B'
		for _, d := range directions {
			dfs(board, i+d[0], j+d[1])
		}
	}

	dfs(board, click[0], click[1])

	return board
}

func count(board [][]byte, i, j int) int {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return 0
	}

	f := func(i, j int) int {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return 0
		}
		if board[i][j] == 'M' {
			return 1
		}

		return 0
	}

	sum := 0
	for _, d := range directions {
		sum += f(i+d[0], j+d[1])
	}

	return sum
}
