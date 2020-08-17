package leetcode

/*
* 110. 平衡二叉树
* https://leetcode-cn.com/problems/balanced-binary-tree/
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var height func(*TreeNode) (bool, int)
	height = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}
		lb, l := height(root.Left)
		rb, r := height(root.Right)

		if !lb || !rb {
			return false, 0
		}

		if v := l - r; v > 1 || v < -1 {
			return false, 0
		}

		return true, max(l, r) + 1
	}

	ok, _ := height(root)

	return ok
}
