package leecode_converter

import "encoding/json"

type TreeConverter string

func (l TreeConverter) Parse() *TreeNode {
	s := string(l)
	var arr []interface{}

	err := json.Unmarshal([]byte(s), &arr)
	if err != nil {
		panic(err)
	}

	if len(arr) == 0 {
		return nil
	}

	var queue []*TreeNode

	for _, v := range arr {
		if v == nil {
			queue = append(queue, nil)
		} else {
			queue = append(queue, &TreeNode{Val: int(v.(float64))})
		}
	}

	for i, j := 0, 1; i < len(queue); i++ {
		if queue[i] == nil {
			continue
		}

		if j < len(queue) {
			queue[i].Left = queue[j]
		}

		if j+1 < len(queue) {
			queue[i].Right = queue[j+1]
		}

		j += 2
	}

	return queue[0]
}
