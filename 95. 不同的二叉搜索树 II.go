package leetcode

/*
* 95. 不同的二叉搜索树 II
* https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
 */

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var f func(int, int) []*TreeNode
	f = func(first, last int) []*TreeNode {
		if first > last {
			return []*TreeNode{nil}
		}

		var allTrees []*TreeNode

		for i := first; i <= last; i++ {
			leftTrees := f(first, i-1)
			rightTrees := f(i+1, last)

			for _, l := range leftTrees {
				for _, r := range rightTrees {
					allTrees = append(allTrees, &TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					})
				}
			}
		}

		return allTrees
	}

	all := f(1, n)

	return all
}
