package leetcode

import "sort"

/*
* 90. 子集 II
* https://leetcode-cn.com/problems/subsets-ii/submissions/
 */

/*
* 难度: 中等
* 通过: true
 */

func subsetsWithDup(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{{}}
	}

	sort.Ints(nums)

	arr := [][]int{
		{nums[0]},
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			arr[len(arr)-1] = append(arr[len(arr)-1], nums[i])
			continue
		}

		arr = append(arr, []int{nums[i]})
	}

	var ans [][]int

	var f func(n int, tmp []int)

	f = func(n int, tmp []int) {
		if n >= len(arr) {
			ans = append(ans, append(tmp[:0:0], tmp...))
			return
		}

		l := len(tmp)
		for i := len(arr[n]); i >= 0; i-- {
			tmp = append(tmp, arr[n][:i]...)
			f(n+1, tmp)
			tmp = tmp[:l]
		}
	}

	f(0, []int{})

	return ans
}
