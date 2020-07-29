package leetcode

import (
	"sort"
	"strings"
)

/*
* 1268. 搜索推荐系统
* https://leetcode-cn.com/problems/search-suggestions-system/
 */

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	ans := make([][]string, len(searchWord)+1)
	ans[0] = products

	for i := 1; i <= len(searchWord); i++ {
		for _, product := range ans[i-1] {
			if strings.HasPrefix(product, searchWord[:i]) {
				ans[i] = append(ans[i], product)
			}
		}
	}

	ans = ans[1:]
	for i := 0; i < len(ans); i++ {
		if len(ans[i]) > 3 {
			ans[i] = ans[i][:3]
		}
	}

	return ans
}
