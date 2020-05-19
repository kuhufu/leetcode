package leetcode

import "math"

/*
* 152. 乘积最大子数组
* https://leetcode-cn.com/problems/maximum-product-subarray/
 */

//又尼玛是动态规划
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mem := make([]int, len(nums)+1)
	mem[0] = 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		tmp := 1
		for i < len(nums) && nums[i] != 0 {
			tmp *= nums[i]
			mem[i+1] = tmp
			i++
		}
	}

	ans := math.MinInt64

	for i := 0; i < len(mem); i++ {
		tmp := mem[i]
		if mem[i] == 0 && i+1 < len(mem) {
			if mem[i+1] == 0 {
				continue
			} else {
				tmp = 1
			}
		}

		for j := i + 1; j < len(mem); j++ {
			if v := mem[j] / tmp; v > ans {
				ans = v
			}

			if mem[j] == 0 {
				break
			}
		}
	}

	return ans
}
