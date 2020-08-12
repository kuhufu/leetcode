package leecode_converter

import "encoding/json"

type GraphConverter string

func (g GraphConverter) Parse() *Node {
	var arr [][]int

	err := json.Unmarshal([]byte(g), &arr)
	if err != nil {
		panic(err)
	}

	nodes := map[int]*Node{
		//1: start,
	}

	for i := 0; i < len(arr); i++ {
		nodes[i+1] = &Node{Val: i + 1}
	}

	for i, neighbors := range arr {
		cur := nodes[i+1]
		for _, neighbor := range neighbors {
			cur.Neighbors = append(cur.Neighbors, nodes[neighbor])
		}
	}

	return nodes[1]
}
