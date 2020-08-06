package leetcode

import (
	"github.com/kuhufu/leetcode/leecode_converter"
)

type ListNode = leecode_converter.ListNode

type TreeNode = leecode_converter.TreeNode

func Slice(s string) leecode_converter.SliceConverter {
	return leecode_converter.SliceConverter(s)
}

func List(s string) leecode_converter.ListConverter {
	return leecode_converter.ListConverter(s)
}

func Tree(s string) leecode_converter.TreeConverter {
	return leecode_converter.TreeConverter(s)
}

func Str(args ...interface{}) string {
	return leecode_converter.Str(args...)
}
