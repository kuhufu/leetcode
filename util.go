package leetcode

import (
	"reflect"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
// 1,2,3,4
// 1, 2, 3, 4
func str2Slice(s string) []int {
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, " ", "", -1)

	items := strings.Split(s, ",")

	var res []int

	for _, item := range items {
		res = append(res, Atoi(item))
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
	default:
		return reflect.DeepEqual(a, b)
	}
}
