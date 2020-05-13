package leetcode

/*
* 102. 二叉树的层序遍历
* https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

/*
* 难度：中等
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int

	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)

		var tmp []int
		for i := 0; i < size; i++ {
			tmp = append(tmp, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		ans = append(ans, tmp)
		q = q[size:]
	}

	return ans
}
