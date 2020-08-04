package leetcode

/*
* 207. 课程表
* https://leetcode-cn.com/problems/course-schedule/
 */

func canFinish(numCourses int, prerequisites [][]int) bool {
	const (
		Init = iota
		Visited
		Ok
	)

	edges := make([][]int, numCourses)
	state := make([]int, numCourses)

	for _, v := range prerequisites {
		edges[v[0]] = append(edges[v[0]], v[1])
	}

	var dfs func(int) bool
	dfs = func(k int) bool {
		if state[k] == Ok {
			return true
		}

		if state[k] == Visited {
			return false
		}

		neighbors := edges[k]
		state[k] = Visited

		for _, neighbor := range neighbors {
			if !dfs(neighbor) {
				return false
			}
		}
		state[k] = Ok
		return true
	}

	for k, _ := range edges {
		if state[k] == Init && !dfs(k) {
			return false
		}
	}

	return true
}
