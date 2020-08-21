package leetcode

import (
	"github.com/kuhufu/leetcode/leecode_converter"
)

type ListNode = leecode_converter.ListNode

type TreeNode = leecode_converter.TreeNode

type Node = leecode_converter.Node

func Slice(s string) leecode_converter.SliceConverter {
	return leecode_converter.SliceConverter(s)
}

func List(s string) *ListNode {
	return leecode_converter.ListConverter(s).Parse()
}

func Tree(s string) *TreeNode {
	return leecode_converter.TreeConverter(s).Parse()
}

func Graph(s string) leecode_converter.GraphConverter {
	return leecode_converter.GraphConverter(s)
}

func Str(args ...interface{}) string {
	return leecode_converter.Str(args...)
}
