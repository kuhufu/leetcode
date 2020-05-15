package leetcode

/*
* 560. 和为K的子数组
* https://leetcode-cn.com/problems/subarray-sum-equals-k/
 */

/*
* 难度：中等
* 技巧：前缀和
 */

//真的很有意思的解法
func subarraySum(nums []int, k int) int {
	ans := 0
	m := map[int]int{
		0: 1,
	}

	sum := 0
	for _, v := range nums {
		sum += v
		ans += m[sum-k]
		m[sum] += 1
	}

	return ans
}
