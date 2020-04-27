package leetcode

/*
* 33. 搜索旋转排序数组
* https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
 */

/*
* 难度: 中等
* 通过: true
 */

//这题画个图，分情况讨论容易理解
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		//判定mid和target在哪个区域
		if nums[m] > target {
			if nums[m] >= nums[l] && target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[m] < nums[l] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

	}

	return -1
}
