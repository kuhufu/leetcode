package cracking_the_coding_interview_6th

/*
* 面试题 08.03. 魔术索引
 */

func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); {
		if i == nums[i] {
			return i
		}

		if i < nums[i] {
			i = nums[i]
		} else {
			i++
		}
	}

	return -1
}
