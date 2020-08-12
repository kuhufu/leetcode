package leetcode

/*
* 133. 克隆图
* https://leetcode-cn.com/problems/clone-graph/
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[*Node]*Node{}

	var dfs func(*Node) *Node
	dfs = func(oldNode *Node) *Node {
		if v, ok := visited[oldNode]; ok {
			return v
		}

		newNode := &Node{Val: oldNode.Val}
		visited[oldNode] = newNode

		for _, node := range oldNode.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(node))
		}

		return newNode
	}

	return dfs(node)
}
