package leetcode

/*
* 136. 只出现一次的数字
* https://leetcode-cn.com/problems/single-number/
* 题解:  https://leetcode-cn.com/problems/single-number/solution/zhi-chu-xian-yi-ci-de-shu-zi-by-leetcode/
 */

/*
* 难度: 简单
* 通过: true
 */

//异或真的很秀
func singleNumber(nums []int) int {
	ans := 0

	for _, v := range nums {
		ans ^= v
	}

	return ans
}
