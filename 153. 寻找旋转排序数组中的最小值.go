package leetcode

/*
* 153. 寻找旋转排序数组中的最小值
* https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
 */

func findMin1(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (hi-lo)/2 + lo
		//if nums[mid] > nums[mid+1] {
		//	return nums[mid+1]
		//}
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		}
	}

	return nums[lo]
}
