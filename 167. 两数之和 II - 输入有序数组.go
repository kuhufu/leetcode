package leetcode

/*
* 167. 两数之和 II - 输入有序数组
* https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
 */

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		v := numbers[l] + numbers[r] //这个解法有点意思
		switch {
		case v > target:
			r--
		case v < target:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}

	return nil
}
