package leetcode

/*
* 130. 被围绕的区域
* https://leetcode-cn.com/problems/surrounded-regions/
 */

func solve(board [][]byte) [][]byte {
	if len(board) == 0 {
		return nil
	}
	m := len(board)
	n := len(board[0])

	for i := 0; i < n; i++ {
		foo(board, 0, i, m, n)
		foo(board, m-1, i, m, n)
	}

	for i := 0; i < m; i++ {
		foo(board, i, 0, m, n)
		foo(board, i, n-1, m, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

	return board
}

func foo(board [][]byte, curX, curY int, m, n int) {
	if curX < 0 || curX >= m || curY < 0 || curY >= n {
		return
	}

	v := board[curX][curY]

	if v == '#' || v == 'X' {
		return
	}

	if v == 'O' {
		board[curX][curY] = '#'
	}

	foo(board, curX+1, curY, m, n)
	foo(board, curX-1, curY, m, n)
	foo(board, curX, curY+1, m, n)
	foo(board, curX, curY-1, m, n)
}
