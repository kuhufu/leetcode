package leetcode

import "fmt"

/*
* 542. 01 矩阵
* https://leetcode-cn.com/problems/01-matrix/
 */

func updateMatrix(matrix [][]int) [][]int {
	dis := Copy(matrix, -1)

	ForEach(matrix, func(i, j int) {
		if matrix[i][j] == 0 {
			dis[i][j] = 0
		}
	})

	for k := 0; ;k++{
		notDone := false
		ForEach(dis, func(i, j int) {
			if dis[i][j] != k {
				return
			}

			b1 := Val(i-1, j, dis, k+1)
			b2 := Val(i+1, j, dis, k+1)
			b3 := Val(i, j-1, dis, k+1)
			b4 := Val(i, j+1, dis, k+1)

			if b1 || b2 || b3 || b4 {
				notDone = true
			}
		})

		if !notDone {
			break
		}
	}

	return dis
}

func ForEach(matrix [][]int, f func(i, j int)) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			f(i, j)
		}
	}
}

func Val(i, j int, dis [][]int, val int) bool {
	if i >= 0 && i < len(dis) && j < len(dis[i]) && j >= 0 {
		if dis[i][j] == -1 {
			dis[i][j] = val
			return true
		}
	}
	return false
}

func Copy(m [][]int, initVal int) [][]int {
	clone := make([][]int, len(m))

	for i := 0; i < len(m); i++ {
		clone[i] = make([]int, len(m[i]))
		for j := 0; j < len(clone[i]); j++ {
			clone[i][j] = initVal
		}
	}

	return clone
}

func Print(matrix [][]int)  {
	for _, a := range matrix {
		fmt.Println(a)
	}
}