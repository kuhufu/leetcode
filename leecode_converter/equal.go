package leecode_converter

import "reflect"

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
	tB := reflect.TypeOf(b)

	if tA != tB {
		return false
	}

	vA := reflect.ValueOf(a)
	vB := reflect.ValueOf(b)

	switch tA.Kind() {
	case reflect.Slice, reflect.Array:
		if vA.Len() != vB.Len() {
			return false
		}
		for i := 0; i < vA.Len(); i++ {
			if !Equal(vA.Index(i).Interface(), vB.Index(i).Interface()) {
				return false
			}
		}
		return true
	}

	switch a.(type) {
	case *ListNode:
		return listEqual(a.(*ListNode), b.(*ListNode))
	case *TreeNode:
		return treeEqual(a.(*TreeNode), b.(*TreeNode))
	case []reflect.Value:
		return reflectValueSliceEqual(a.([]reflect.Value), b.([]reflect.Value))
	default:
		return reflect.DeepEqual(a, b)
	}
}

func treeEqual(s, t *TreeNode) bool {
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

	return treeEqual(s.Left, t.Left) && treeEqual(s.Right, t.Right)
}

func reflectValueSliceEqual(a, b []reflect.Value) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !reflect.DeepEqual(a[i].Interface(), b[i].Interface()) {
			return false
		}
	}

	return true
}
