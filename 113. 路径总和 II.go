package leetcode

/*
* 113. 路径总和 II
* https://leetcode-cn.com/problems/path-sum-ii/
 */

func pathSum(root *TreeNode, sum int) [][]int {
	var ans [][]int

	var f func(root *TreeNode, stack []int, sum int)
	f = func(root *TreeNode, stack []int, sum int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
			ans = append(ans, append(stack[:0:0], append(stack, root.Val)...))
		}

		f(root.Left, append(stack, root.Val), sum-root.Val)
		f(root.Right, append(stack, root.Val), sum-root.Val)
	}

	f(root, nil, sum)

	return ans
}
