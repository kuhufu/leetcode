package leetcode

import (
	"math"
)

/*
* 23. 合并K个排序链表
* https://leetcode-cn.com/problems/merge-k-sorted-lists/
 */

//第一种做法，遍历所有链表的第一个Node找到最小的，重复这个过程
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	ans := &ListNode{}
	start := ans
	pos := make([]*ListNode, len(lists))

	for i, _ := range lists {
		pos[i] = lists[i]
	}

	Init := &ListNode{Val: math.MaxInt64}
	for {
		min := Init
		minPos := -1
		for i, node := range pos {
			if node != nil {
				if node.Val < min.Val {
					min = node
					minPos = i
				}
			}
		}

		if min == Init {
			break
		}

		pos[minPos] = min.Next

		start.Next = min
		start = min
	}

	start.Next = nil

	return ans.Next
}

//第二种做法，因为链表是有序的，所以可以直接像归并排序那样，进行二路归并，时间复杂度远低于第一种
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 0 {
		return nil
	}

	mid := len(lists) / 2
	l1 := mergeKLists2(lists[:mid])
	l2 := mergeKLists2(lists[mid:])

	ans := &ListNode{}
	start := ans

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			start.Next = l1
			l1 = l1.Next
		} else {
			start.Next = l2
			l2 = l2.Next
		}
		start = start.Next
	}

	if l1 == nil {
		start.Next = l2
	}

	if l2 == nil {
		start.Next = l1
	}

	return ans.Next
}
