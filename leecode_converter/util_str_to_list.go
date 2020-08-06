package leecode_converter

import (
	"strconv"
	"strings"
)

type ListConverter string

func (l ListConverter) Parse() *ListNode {
	s := string(l)
	res := &ListNode{}
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	items := strings.Split(s, "->")

	cur := res
	for _, item := range items {
		v, _ := strconv.Atoi(item)
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return res.Next
}
