package leetcode

import "math"

/*
* 124. 二叉树中的最大路径和
* https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
 */

/*
* 难度: 困难
 */

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64

	var f func(*TreeNode) int

	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		var curVal = root.Val
		var curMax = root.Val
		var cul = root.Val
		var ret = root.Val

		if root.Left != nil {
			l := f(root.Left)
			cul += l
			curMax = max(curMax, l)
			ret = max(ret, curVal+l)
		}

		if root.Right != nil {
			r := f(root.Right)
			cul += r
			curMax = max(curMax, r)
			ret = max(ret, curVal+r)
		}

		if curMax = max(curMax, cul); curMax > ans {
			ans = curMax
		}

		return ret
	}

	if v := f(root); v > ans {
		ans = v
	}

	return ans
}
