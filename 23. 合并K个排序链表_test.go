package leetcode

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	Run(t, mergeKLists2, []Test{
		{
			Args{
				[]*ListNode{
					List("1->4->5"),
					List("1->3->4"),
				},
			},
			Want{List("1->1->3->4->4->5")},
		},
		{
			Args{
				[]*ListNode{
					List("1->4->5"),
					List("1->3->4"),
					List("2->6"),
				},
			},
			Want{List("1->1->2->3->4->4->5->6")},
		},
	})
}
