package leetcode

/*
* 1004. 最大连续1的个数 III
* https://leetcode-cn.com/problems/max-consecutive-ones-iii/
 */

//滑动窗口
func longestOnes(A []int, K int) int {
	ans := 0

	right := 0
	left := 0
	k := 0

	for right < len(A) {
		if A[right] == 0 {
			k++
		}

		right++

		for k > K {
			if A[left] == 0 {
				k--
			}
			left++
		}

		if v := right-left; v > ans {
			ans = v
		}
	}

	return ans
}
