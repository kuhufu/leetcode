package leetcode

import "sort"

/*
* 16. 最接近的三数之和
* https://leetcode-cn.com/problems/3sum-closest/
 */

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	ans := 0
	dis := 1 << 61

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			curDis := sum - target
			if Abs(curDis) < Abs(dis) {
				ans = sum
				dis = curDis
			}
			switch {
			case curDis > 0:
				k--
			case curDis < 0:
				j++
			case curDis == 0:
				return sum
			}
		}
	}

	return ans
}
