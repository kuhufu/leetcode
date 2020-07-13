package leetcode

/*
* 350. 两个数组的交集 II
* https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
 */

func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	c := map[int]int{}

	//选择长度小的数组做map
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for _, num := range nums2 {
		c[num] += 1
	}

	for _, num := range nums1 {
		if v, ok := c[num]; ok && v > 0 {
			c[num] = v - 1
			ans = append(ans, num)
		}
	}

	return ans
}
