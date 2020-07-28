package leetcode

/*
* 104. 二叉树的最大深度
* https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
