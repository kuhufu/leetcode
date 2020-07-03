package leetcode

/*
* 55. 跳跃游戏
* https://leetcode-cn.com/problems/jump-game/
 */

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	maxAccess := 0

	endIdx := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		curMaxAccess := nums[i] + i
		
		if curMaxAccess >= endIdx {
			return true
		}

		if i == maxAccess && curMaxAccess == maxAccess {
			return false
		}

		if curMaxAccess > maxAccess {
			maxAccess = curMaxAccess
		}
	}

	return false
}
