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
func Str2Slice(s string) []int {
	var res []int
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		panic(err)
	}

	return res
}

//input format:
// [[1,2,3,4],[1,2,3,4]]
func Str22DSlice(s string) [][]int {
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
	case []reflect.Value:
		return reflectValueSliceEqual(a.([]reflect.Value), b.([]reflect.Value))
	default:
		return reflect.DeepEqual(a, b)
	}
}

//format:
//[1,2,3,4],层序遍历？的结果，如果出现null，这表示该节点为空
func Str2Tree(s string) *TreeNode {
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

func treeEqual(s, t *TreeNode) bool {
	switch {
	case s == nil && t == nil:
		return true
	case s != nil && t == nil:
		return false
	case s == nil && t != nil:
		return false
	case s.Val != t.Val:
		return false
	}

	return treeEqual(s.Left, t.Left) && treeEqual(s.Right, t.Right)
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
				continue
			}

			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}

		queue = queue[size:]
	}

	//去掉末尾的 nil
	var i int
	for i = len(nodes) - 1; i >= 0; i-- {
		if nodes[i] != nil {
			break
		}
	}

	nodes = nodes[:i+1]

	bytes, err := json.Marshal(&nodes)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func Str(args ...interface{}) string {
	var res []string
	for _, arg := range args {
		res = append(res, str(arg))
	}

	return strings.Join(res, ", ")
}

func str(v interface{}) string {
	switch v := v.(type) {
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

func reflectValueSliceEqual(a, b []reflect.Value) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !reflect.DeepEqual(a[i].Interface(), b[i].Interface()) {
			return false
		}
	}

	return true
}
