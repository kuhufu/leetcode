package leetcode

import "sort"

/*
* 56. 合并区间
* https://leetcode-cn.com/problems/merge-intervals/
 */

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	j := 0

	for i := 1; i < len(intervals); i++ {
		curInterval := intervals[i]
		curAnsInterval := ans[j]

		if curInterval[0] > curAnsInterval[1] {
			ans = append(ans, curInterval)
			j++
			continue
		}

		if curInterval[1] > curAnsInterval[1] {
			curAnsInterval[1] = curInterval[1]
			continue
		}
	}

	return ans
}
