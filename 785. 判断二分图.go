package leetcode

/*
* 785. 判断二分图
* https://leetcode-cn.com/problems/is-graph-bipartite/
 */

const (
	UNCOLORED = iota
	RED
	BLUE
)

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if color[i] == UNCOLORED {
			color[i] = RED
		}

		var curC = color[i] //当前节点颜色
		var nextC int       //相邻节点的颜色
		queue := []int{i}

		for len(queue) != 0 {
			switch curC { //根据当前节点颜色决定相邻节点颜色
			case RED:
				nextC = BLUE
			case BLUE:
				nextC = RED
			}

			size := len(queue)
			for _, node := range queue[:size] {
				for _, v := range graph[node] {
					switch color[v] {
					case UNCOLORED: //相邻节点未染色
						color[v] = nextC         //染色
						queue = append(queue, v) //加入队列
					case curC: //相邻节点颜色跟当前节点颜色相同
						return false //不满足条件
					}
				}
			}
			queue = queue[size:]
			curC = nextC
		}
	}

	return true
}
