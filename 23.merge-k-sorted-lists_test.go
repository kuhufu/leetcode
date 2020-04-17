package leetcode

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{
			[]*ListNode{
				str2List("1->4->5"),
				str2List("1->3->4"),
			},
			str2List("1->1->3->4->4->5"),
		},
		{
			[]*ListNode{
				str2List("1->4->5"),
				str2List("1->3->4"),
				str2List("2->6"),
			},
			str2List("1->1->2->3->4->4->5->6"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(listArr2Str(test.lists), func(t *testing.T) {
			res := mergeKLists2(test.lists)

			if !Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
