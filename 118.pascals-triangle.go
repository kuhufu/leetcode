package leetcode

/*
* 118. 杨辉三角
* https://leetcode-cn.com/problems/pascals-triangle/
 */

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			if j-1 < 0 || j >= i || i-1 < 0 {
				ans[i][j] = 1
				continue
			}

			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
