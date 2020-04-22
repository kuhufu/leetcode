package leetcode

/*
* 1248. 统计「优美子数组」
* https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
 */

//数学法还能理解
func numberOfSubarrays(nums []int, k int) int {
	ans := 0

	pos := []int{-1}
	for i, num := range nums {
		if num%2 == 1 {
			pos = append(pos, i)
		}
	}
	pos = append(pos, len(nums))

	for i := 1; i+k < len(pos); i++ {
		left := pos[i] - pos[i-1]
		right := pos[i+k] - pos[i+k-1]

		ans += left * right
	}

	return ans
}

//这个解法，看了官方题解也不懂
func numberOfSubarrays2(nums []int, k int) int {
	cnt := make([]int, len(nums)+1)

	odd := 0
	ans := 0

	cnt[0] = 1

	for i := 0; i < len(nums); i++ {
		odd += nums[i] & 1

		if odd >= k {
			ans += cnt[odd-k]
		}

		cnt[odd] += 1
	}

	return ans
}
