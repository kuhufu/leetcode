package leetcode

/*
* 108. 将有序数组转换为二叉搜索树
 */

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[0:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return root
}
