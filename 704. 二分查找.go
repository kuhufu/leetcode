package leetcode

/*
* 704. 二分查找
* https://leetcode-cn.com/problems/binary-search/
 */

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid
		} else {
			return mid
		}
	}

	return -1
}

func binarySearch2(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

//go搜索标准库返回的是索引是target的插入位置 索引 >= 0 不会小于0
// 设返回的索引为 i, target 为要搜索的值
// 1. i < len(nums) && nums[i] == target
// 2. i == len(nums) 说明 target 应该添加到数组尾部
func golangStdBinarySearch1(nums []int, target int) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func golangStdBinarySearch2(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}
