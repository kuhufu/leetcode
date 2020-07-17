package leetcode

/*
* 35. 搜索插入位置
* https://leetcode-cn.com/problems/search-insert-position/
 */

/*
* 难度: 简单
 */

func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums)

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
