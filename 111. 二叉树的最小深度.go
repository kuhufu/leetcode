package leetcode

/*
* 111. 二叉树的最小深度
* https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
 */

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 || r == 0 {
		return 1 + l + r
	}

	return min(l, r) + 1
}
