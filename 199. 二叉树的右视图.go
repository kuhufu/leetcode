package leetcode

/*
* 199. 二叉树的右视图
* https://leetcode-cn.com/problems/binary-tree-right-side-view/
 */

//层序遍历
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		ans = append(ans, queue[0].Val)

		for i := 0; i < size; i++ {
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
		}
		queue = queue[size:]
	}

	return ans
}
