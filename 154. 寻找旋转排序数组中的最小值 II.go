package leetcode

/*
* 154. 寻找旋转排序数组中的最小值 II
* https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
 */

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi--
		}
	}
	return nums[lo]
}
