package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* 572. 另一个树的子树
* https://leetcode-cn.com/problems/subtree-of-another-tree/
 */

/*
* 难度: 简单
 */

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s != nil && s.Val == t.Val {
		if treeEqual1(s, t) {
			return true
		}
	}

	if s == nil {
		return false
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func treeEqual1(s, t *TreeNode) bool {
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

	return treeEqual1(s.Left, t.Left) && treeEqual1(s.Right, t.Right)
}

func isSubtree2(s *TreeNode, t *TreeNode) bool {
	sp := prefixTraversal(s)
	tp := prefixTraversal(t)

	fmt.Println(sp)
	fmt.Println(tp)

	return strings.Index(sp, tp) != -1
}

func prefixTraversal(t *TreeNode) string {
	ans := strings.Builder{}

	var f func(*TreeNode)
	f = func(t *TreeNode) {
		ans.WriteString(strconv.Itoa(t.Val))
		if t.Left == nil {
			ans.Write([]byte("l("))
			ans.WriteString(strconv.Itoa(t.Val))
			ans.Write([]byte(")"))
		}

		if t.Right == nil {
			ans.Write([]byte("r("))
			ans.WriteString(strconv.Itoa(t.Val))
			ans.Write([]byte(")"))
		}

		if t.Left != nil {
			f(t.Left)
		}

		if t.Right != nil {
			f(t.Right)
		}
	}

	f(t)

	return ans.String()
}
