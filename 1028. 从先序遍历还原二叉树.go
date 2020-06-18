package leetcode

/*
* 1028. 从先序遍历还原二叉树
 */

func recoverFromPreorder(S string) *TreeNode {
	var dfs func(int) *TreeNode
	dfs = func(curLevel int) *TreeNode {
		if S == "" {
			return nil
		}

		level := 0
		for i := 0; i < len(S) && S[i] == '-'; i++ {
			level++
		}

		if level != curLevel {
			return nil
		}

		S = S[level:]

		value := 0
		for len(S) != 0 && S[0] >= '0' && S[0] <= '9' {
			value = value*10 + int(S[0]-'0')
			S = S[1:]
		}

		return &TreeNode{
			Val:   value,
			Left:  dfs(level + 1),
			Right: dfs(level + 1),
		}
	}

	return dfs(0)
}
