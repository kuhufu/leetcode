package leetcode

/*
* 845. 数组中的最长山脉
* https://leetcode-cn.com/problems/longest-mountain-in-array/
 */

/*
* 难度: 中等
* 通过: 通过
 */

func longestMountain(A []int) int {
	ans := 0

	for i := 1; i+1 < len(A); i++ {
		if A[i-1] >= A[i] || A[i+1] >= A[i] {
			continue
		}

		tmpMax := 1

		var m int
		for m = i; m-1 >= 0 && A[m-1] < A[m]; m-- {
			tmpMax++
		}

		for m = i; m+1 < len(A) && A[m+1] < A[m]; m++ {
			tmpMax++
		}

		if tmpMax > ans {
			ans = tmpMax
		}

		i = m
	}

	return ans
}
