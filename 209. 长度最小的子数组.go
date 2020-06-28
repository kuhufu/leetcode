package leetcode

/*
* 209. 长度最小的子数组
* https://leetcode-cn.com/problems/minimum-size-subarray-sum/submissions/
 */

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := len(nums) + 1
	sum := 0

	i := -1
	j := 0
	for j < len(nums) {
		sum += nums[j]
		if sum >= s {
			if v := j - i; v < ans {
				ans = v
			}

			i++
			sum -= nums[i] + nums[j]
			continue
		}
		j++
	}

	if ans == len(nums)+1 {
		return 0
	}

	return ans
}
