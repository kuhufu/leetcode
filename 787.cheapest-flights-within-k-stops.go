package leetcode

/*
* 787. K 站中转内最便宜的航班
* https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
 */

//这题我感觉自己的脑子坏掉了
//没有剪枝是不行的，第一个测试用例会超时
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	edges := map[int]map[int]int{}
	for _, f := range flights {
		if _, ok := edges[f[0]]; !ok {
			edges[f[0]] = map[int]int{}
		}

		if _, ok := edges[f[1]]; !ok {
			edges[f[1]] = map[int]int{}
		}

		edges[f[0]][f[1]] = f[2]
	}

	queue := [][]int{{src, 0}}

	ans := 1 << 30

	K += 1
	for len(queue) != 0 && K >= 0 {

		size := len(queue)

		for i := 0; i < size; i++ {
			item := queue[0]
			queue = queue[1:]

			if item[0] == dst && item[1] < ans {
				ans = item[1]
			}

			neighbors := edges[item[0]]

			for neighbor, cost := range neighbors {
				if K >= 0 {
					queue = append(queue, []int{neighbor, cost + item[1]})
				}
			}
		}
		K--
	}

	return ans
}
