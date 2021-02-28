package leetcode

/*
* 896. 单调数列
* https://leetcode-cn.com/problems/monotonic-array/
 */

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	dir := "DESC"
	if A[len(A)-1] > A[0] {
		dir = "ASC"
	}

	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > A[i-1] && dir != "ASC":
			return false
		case A[i] < A[i-1] && dir != "DESC":
			return false
		}
	}

	return true
}
