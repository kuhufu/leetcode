package cracking_the_coding_interview_6th

/*
* 面试题 16.11. 跳水板
* https://leetcode-cn.com/problems/diving-board-lcci/
 */

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}

	if shorter == longer {
		return []int{shorter * k}
	}

	ans := make([]int, 0, k+1)
	for i := 0; i <= k; i++ {
		ans = append(ans, longer*i+shorter*(k-i))
	}

	return ans
}
