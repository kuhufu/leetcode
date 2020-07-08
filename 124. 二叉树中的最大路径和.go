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
		var curMax = root.Val //（左子树+当前节点），（当前节点），（当前节点+右子树），（当前节点+右子树+左子树）四者之间的最大值
		var sum = root.Val    //（左子树+右子树+当前节点）的和
		var ret = root.Val    //（左子树+当前节点），（当前节点），（当前节点+右子树）三者之间的最大值

		if root.Left != nil {
			l := f(root.Left)
			sum += l
			curMax = max(curMax, l)
			ret = max(ret, curVal+l)
		}

		if root.Right != nil {
			r := f(root.Right)
			sum += r
			curMax = max(curMax, r)
			ret = max(ret, curVal+r)
		}

		if curMax = max(curMax, max(ret, sum)); curMax > ans {
			ans = curMax
		}

		return ret
	}

	return ans
}
