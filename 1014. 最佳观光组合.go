package leetcode

/*
* 1014. 最佳观光组合
* https://leetcode-cn.com/problems/best-sightseeing-pair/
 */

//这题的解法很有意思值得一看
func maxScoreSightseeingPair(A []int) int {
	ans := 0
	mx := A[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(A); i++ {
		ans = max(A[i]-i+mx, ans)
		mx = max(A[i]+i, mx)
	}

	return ans
}
