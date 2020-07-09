package coding_interviews_2th

import (
	"sort"
)

/*
* 面试题51. 数组中的逆序对
* 难度: 困难
* https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
 */

/*
* 提交通过: false
 */

//题解都看不懂，难受
func reversePairs(nums []int) int {
	n := len(nums)
	tmp := append(nums[:0:0], nums...)
	sort.Ints(tmp)

	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}
