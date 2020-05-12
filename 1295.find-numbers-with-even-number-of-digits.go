package leetcode

/*
* 1295. 统计位数为偶数的数字
* https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
 */

/*
* 难度: 过于简单
 */

func findNumbers(nums []int) int {
	ans := 0

	for _, num := range nums {
		count := 0
		for num != 0 {
			num /= 10
			count++
		}

		if count%2 == 0 {
			ans++
		}
	}

	return ans
}
