package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := ListLen(l1)
	n2 := ListLen(l2)

	dis := n2 - n1
	if dis >= 0 {
		l1, l2 = l2, l1
	} else {
		dis = -dis
	}

	node, i := dfs(dis, l1, l2)
	if i != 0 {
		return &ListNode{Val: 1, Next: node}
	} else {
		return node
	}
}

func dfs(dis int, l1 *ListNode, l2 *ListNode) (*ListNode, int) {
	if l1 == nil  {
		return nil, 0
	}

	var (
		v        int
		nextNode *ListNode
		carryBit int
	)

	if dis > 0 {
		nextNode, carryBit = dfs(dis-1, l1.Next, l2)
		v = l1.Val + 0 + carryBit
	} else {
		nextNode, carryBit = dfs(dis-1, l1.Next, l2.Next)
		v = l1.Val + l2.Val + carryBit
	}

	return &ListNode{Val: v % 10, Next: nextNode}, v / 10
}

func ListLen(l *ListNode) int {
	count := 0
	for l := l; l != nil; l = l.Next {
		count++
	}

	return count
}
