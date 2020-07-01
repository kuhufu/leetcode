package leetcode

/*
* 718. 最长重复子数组
 */

/*
* tag: 滑动窗口
 */

func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	ret := 0

	for i := 0; i < n; i++ {
		len := min(m, n-i)
		if len <= ret { //小优化
			break
		}
		maxLen := maxLength(A, B, i, 0, len)
		ret = max(ret, maxLen)
	}
	for i := 0; i < m; i++ {
		len := min(n, m-i)
		if len <= ret { //小优化
			break
		}
		maxLen := maxLength(A, B, 0, i, len)
		ret = max(ret, maxLen)
	}
	return ret
}

func maxLength(A, B []int, addA, addB, len int) int {
	ret, k := 0, 0
	for i := 0; i < len; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}
