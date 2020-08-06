package leetcode

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	Run(t, mergeKLists2, []Test{
		{
			Args{
				[]*ListNode{
					List("1->4->5").Parse(),
					List("1->3->4").Parse(),
				},
			},
			Want{List("1->1->3->4->4->5").Parse()},
		},
		{
			Args{
				[]*ListNode{
					List("1->4->5").Parse(),
					List("1->3->4").Parse(),
					List("2->6").Parse(),
				},
			},
			Want{List("1->1->2->3->4->4->5->6").Parse()},
		},
	})
}
