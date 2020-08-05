package leetcode

/*
* 337. 打家劫舍 III
* https://leetcode-cn.com/problems/house-robber-iii/
 */

func robIII(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (rob, notRob int) {
		if root == nil {
			return 0, 0
		}

		lRob, lNotRob := dfs(root.Left)
		rRob, rNotRob := dfs(root.Right)

		rob = root.Val + lNotRob + rNotRob
		notRob = max(lRob, lNotRob) + max(rRob, rNotRob)

		return
	}

	rob, notRob := dfs(root)

	return max(rob, notRob)
}
