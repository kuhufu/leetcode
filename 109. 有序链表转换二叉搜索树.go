package leetcode

/*
* 109. 有序链表转换二叉搜索树
* https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
 */

func sortedListToBST(head *ListNode) *TreeNode {
	length := getLen(head)

	root := buildTree(&head, 0, length-1)
	return root
}

func buildTree(list **ListNode, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	root := &TreeNode{}
	mid := (hi-lo)/2 + lo
	root.Left = buildTree(list, lo, mid-1)

	root.Val = (*list).Val
	*list = (*list).Next

	root.Right = buildTree(list, mid+1, hi)

	return root
}

func getLen(head *ListNode) int {
	ans := 0

	for head != nil {
		ans++
		head = head.Next
	}

	return ans
}
