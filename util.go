package leetcode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

//input format: 1->2->3->4->5
func str2List(s string) *ListNode {
	res := &ListNode{}
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	items := strings.Split(s, "->")

	cur := res
	for _, item := range items {
		cur.Next = &ListNode{Val: Atoi(item)}
		cur = cur.Next
	}

	return res.Next
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

func listArr2Str(arr []*ListNode) string {
	var strArr []string

	for _, item := range arr {
		strArr = append(strArr, list2Str(item))
	}

	return "[" + strings.Join(strArr, ",") + "]"
}

//input format:
// [1,2,3,4]
func str2Slice(s string) []int {
	var res []int
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		panic(err)
	}

	return res
}

//input format:
// [[1,2,3,4],[1,2,3,4]]
func str22DSlice(s string) [][]int {
	var res [][]int
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		panic(err)
	}

	return res
}

func slice2Str(a []int) string {
	var strArr []string
	for _, v := range a {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return "[" + strings.Join(strArr, ",") + "]"
}

//有环链表不适用
func listEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 == nil && l1 == nil {
		return true
	}

	return false
}

func Equal(a, b interface{}) bool {
	tA := reflect.TypeOf(a)
	Tb := reflect.TypeOf(b)

	if tA != Tb {
		return false
	}

	switch a.(type) {
	case *ListNode:
		return listEqual(a.(*ListNode), b.(*ListNode))
	case *TreeNode:
		return treeEqual(a.(*TreeNode), b.(*TreeNode))
	default:
		return reflect.DeepEqual(a, b)
	}
}

//format:
//[1,2,3,4],层序遍历的结果，如果出现null，这表示该节点为空
func str2Tree(s string) *TreeNode {
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

	for i := 0; i < len(queue); i++ {
		if queue[i] == nil {
			continue
		}

		if i*2+1 >= len(queue) {
			break
		}
		queue[i].Left = queue[i*2+1]

		if i*2+2 >= len(queue) {
			break
		}
		queue[i].Right = queue[i*2+2]
	}

	return queue[0]
}

func treeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 != nil && t2 == nil || t1 == nil && t2 != nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	left := treeEqual(t1.Left, t2.Left)
	if !left {
		return false
	}

	right := treeEqual(t1.Right, t2.Right)
	return right
}

func tree2Str(root *TreeNode) string {
	queue := []*TreeNode{root}
	var nodes []interface{}

	for len(queue) != 0 {
		size := len(queue)

		var tmpArr []interface{}

		for _, v := range queue {
			if v == nil {
				tmpArr = append(tmpArr, nil)
			} else {
				tmpArr = append(tmpArr, v.Val)
			}
		}

		nodes = append(nodes, tmpArr...)
		for i := 0; i < size; i++ {
			if queue[i] == nil {
				queue = append(queue, nil)
				queue = append(queue, nil)
				continue
			}

			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}

		queue = queue[size:]

		//去掉末尾的 nil
		var i int
		for i = len(queue) - 1; i >= 0; i-- {
			if queue[i] != nil {
				break
			}
		}

		queue = queue[:i+1]
	}

	bytes, err := json.Marshal(&nodes)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func Str(v interface{}) string {
	switch v := v.(type) {
	case *TreeNode:
		return tree2Str(v)
	case *ListNode:
		return list2Str(v)
	case []*ListNode:
		return listArr2Str(v)
	default:
		return fmt.Sprint(v)
	}
}
