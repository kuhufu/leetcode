package leetcode

/*
* 221. 最大正方形
* https://leetcode-cn.com/problems/maximal-square/
 */

/*
* 难度: 中等
* 又是动态规划，建议看题解
 */

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	M := len(matrix)
	N := len(matrix[0])

	ans := byte('0')

	for i := 0; i < N; i++ {
		if matrix[0][i] == '1' {
			ans = '1'
		}
	}

	if ans == '0' {
		for i := 0; i < M; i++ {
			if matrix[i][0] == '1' {
				ans = '1'
			}
		}
	}

	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if matrix[i][j] == '1' {
				matrix[i][j] = MinOf(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]) + 1
				if matrix[i][j] > ans {
					ans = matrix[i][j]
				}
			}
		}
	}

	ans -= '0'

	return int(ans * ans)
}

func MinOf(a, b, c byte) byte {
	if b < a {
		if c < b {
			return c
		} else {
			return b
		}
	} else {
		if c < a {
			return c
		} else {
			return a
		}
	}
}
