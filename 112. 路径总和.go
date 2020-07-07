package leetcode

/*
* 112. 路径总和
 */

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil { //左右节点都是nil，当前节点才是叶子节点
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
