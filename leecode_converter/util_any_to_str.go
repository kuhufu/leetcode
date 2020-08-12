package leecode_converter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func slice2Str(a []int) string {
	var strArr []string
	for _, v := range a {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return "[" + strings.Join(strArr, ",") + "]"
}

// return format: 1->2->3
func list2Str(l *ListNode) string {
	var vals []string
	for l != nil {
		vals = append(vals, strconv.Itoa(l.Val))
		l = l.Next
	}

	return strings.Join(vals, "->")
}

func graph2Str(node *Node) string {
	visited := map[int]*Node{}

	var dfs func(*Node)
	dfs = func(node *Node) {
		if _, ok := visited[node.Val]; ok {
			return
		}

		visited[node.Val] = node

		for _, neighbor := range node.Neighbors {
			dfs(neighbor)
		}
	}

	dfs(node)

	arr := make([][]int, len(visited))

	for i := 0; i < len(arr); i++ {
		for _, neighbor := range visited[i+1].Neighbors {
			arr[i] = append(arr[i], neighbor.Val)
		}
	}

	bytes, _ := json.Marshal(arr)
	return string(bytes)
}

func listArr2Str(arr []*ListNode) string {
	var strArr []string

	for _, item := range arr {
		strArr = append(strArr, list2Str(item))
	}

	return "[" + strings.Join(strArr, ",") + "]"
}

func tree2Str(root *TreeNode) string {
	queue := []*TreeNode{root}

	var strSlice []string
	for len(queue) != 0 {
		tmp := queue[0]
		if tmp != nil {
			queue = append(queue, tmp.Left, tmp.Right)
			strSlice = append(strSlice, strconv.Itoa(tmp.Val))
		} else {
			strSlice = append(strSlice, "null")
		}
		queue = queue[1:]
	}

	//去除尾部null
	for i := len(strSlice) - 1; i >= 0 && strSlice[i] == "null"; i-- {
		strSlice = strSlice[:i]
	}

	return "[" + strings.Join(strSlice, ",") + "]"
}

func str(v interface{}) string {
	switch v := v.(type) {
	case *Node:
		return graph2Str(v)
	case *TreeNode:
		return tree2Str(v)
	case *ListNode:
		return list2Str(v)
	case []*ListNode:
		return listArr2Str(v)
	case string:
		return v
	case []byte, [][]byte:
		return fmt.Sprint(v)
	default:
		bytes, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		return string(bytes)
	}
}

func Str(args ...interface{}) string {
	var res []string
	for _, arg := range args {
		res = append(res, str(arg))
	}

	return strings.Join(res, ", ")
}
