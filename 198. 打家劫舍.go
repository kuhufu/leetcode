package leetcode

/*
* 198. 打家劫舍
* https://leetcode-cn.com/problems/house-robber/
 */

func rob(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}

	if N == 1 {
		return nums[0]
	}

	if N == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])
	nums[2] = max(nums[2]+nums[0], nums[1])

	for i := 3; i < N; i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}

	return nums[N-1]
}

func rob2(nums []int) int {
	first := 0
	second := 0

	for _, num := range nums {
		temp := second
		second = max(second, first+num)
		first = temp
	}
	return second
}
