package leetcode

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	Run(t, mergeKLists2, []Test{
		{
			Args{
				[]*ListNode{
					str2List("1->4->5"),
					str2List("1->3->4"),
				},
			},
			Want{str2List("1->1->3->4->4->5")},
		},
		{
			Args{
				[]*ListNode{
					str2List("1->4->5"),
					str2List("1->3->4"),
					str2List("2->6"),
				},
			},
			Want{str2List("1->1->2->3->4->4->5->6")},
		},
	})
}
