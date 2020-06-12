package leetcode

import (
	"sort"
)

/*
* 15. 三数之和
* https://leetcode-cn.com/problems/3sum/
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int

	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
		if k-1 >= 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			switch {
			case sum < 0:
				i++
			case sum > 0:
				j--
			case sum == 0:
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			}
		}
	}

	return ans
}
