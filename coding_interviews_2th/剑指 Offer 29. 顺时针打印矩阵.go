package coding_interviews_2th

/*
* 面试题29. 顺时针打印矩阵
* https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}

	u := 0
	d := len(matrix) - 1

	l := 0
	r := len(matrix[0]) - 1

	var ans []int

	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[u][i])
		}

		if u++; u > d {
			break
		}

		for i := u; i <= d; i++ {
			ans = append(ans, matrix[i][r])
		}

		if r--; r < l {
			break
		}

		for i := r; i >= l; i-- {
			ans = append(ans, matrix[d][i])
		}

		if d--; d < u {
			break
		}

		for i := d; i >= u; i-- {
			ans = append(ans, matrix[i][l])
		}

		if l++; l > r {
			break
		}
	}

	return ans
}
