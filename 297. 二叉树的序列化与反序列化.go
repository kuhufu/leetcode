package leetcode

import (
	"encoding/json"
	"strconv"
	"strings"
)

/*
* 297. 二叉树的序列化与反序列化
* https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
 */

/*
* 这与 LeetCode 目前使用的方式一致: https://support.leetcode-cn.com/hc/kb/article/1194353/
 */

type Codec struct {
}

func TreeCodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
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

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var arr []interface{}
	s := data

	err := json.Unmarshal([]byte(s), &arr)
	if err != nil {
		panic(err)
	}

	if len(arr) == 0 {
		return nil
	}

	var queue []*TreeNode

	for _, item := range arr {
		var v *TreeNode
		if item != nil {
			v = &TreeNode{Val: int(item.(float64))}
		}
		queue = append(queue, v)
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

	if len(queue) == 0 {
		return nil
	}
	return queue[0]
}
