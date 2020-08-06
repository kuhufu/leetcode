package leetcode

import (
	"math"
)

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

		var val = root.Val
		var sum = root.Val // 左子树+右子树+当前节点的和
		var ret = root.Val // [左子树+当前节点, 当前节点, 当前节点+右子树]三者之间的最大值

		if root.Left != nil {
			l := f(root.Left)
			sum += l
			ret = max(ret, val+l)
		}

		if root.Right != nil {
			r := f(root.Right)
			sum += r
			ret = max(ret, val+r)
		}

		// [左子树+当前节点, 当前节点, 当前节点+右子树, 当前节点+左子树+右子树] 四者之间的最大值
		// ret 的值是 [左子树+当前节点, 当前节点, 当前节点+右子树]三者之间的最大值
		// sum 的值是 当前节点+左子树+右子树 的和
		// 所以:
		if v := max(ret, sum); v > ans {
			ans = v
		}

		return ret
	}

	f(root)

	return ans
}
