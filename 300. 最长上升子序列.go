package leetcode

/*
* 300. 最长上升子序列
* https://leetcode-cn.com/problems/longest-increasing-subsequence/
 */

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	ans := 1
	for i := 1; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > max {
				max = dp[j]
			}
		}

		dp[i] = max + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
