package leetcode

/*
* 378. 有序矩阵中第K小的元素
 */

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)/2 //正常写法是 mid=(lo+hi)/2，这样写是避免溢出
		if gteK(matrix, mid, k) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

//判断mid的最大排名是否大等于k
func gteK(matrix [][]int, mid, k int) bool {
	n := len(matrix)

	i, j := n-1, 0
	count := 0
	for i >= 0 {
		count += j
		for j < n && matrix[i][j] <= mid {
			count++
			j++
		}
		i--
	}

	return count >= k
}
