package leetcode

import "math/rand"

/*
* 215. 数组中的第K个最大元素
* https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 */

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, lo, hi, k int) int {

	for {
		mid := lo + partition(nums[lo:hi+1])
		switch {
		case mid+1 == k:
			return nums[mid]
		case mid+1 < k:
			lo = mid + 1
		case mid+1 > k:
			hi = mid - 1
		}
	}
}

func partition(nums []int) int {
	idx := rand.Intn(len(nums))
	nums[0], nums[idx] = nums[idx], nums[0]

	k := 1
	v := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] >= v {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	nums[0], nums[k-1] = nums[k-1], nums[0]

	return k - 1
}
