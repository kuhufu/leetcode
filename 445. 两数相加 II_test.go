package leetcode

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbersII(t *testing.T) {
	//(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	node := addTwoNumbers(l1, l2)

	fmt.Println(node)
}
