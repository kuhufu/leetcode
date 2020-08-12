package leecode_converter

import (
	"fmt"
	"testing"
)

func Slice(s string) SliceConverter {
	return SliceConverter(s)
}

func List(s string) ListConverter {
	return ListConverter(s)
}

func Tree(s string) TreeConverter {
	return TreeConverter(s)
}

func Graph(s string) GraphConverter {
	return GraphConverter(s)
}

func Test_List(t *testing.T) {

	list := List("1->2->3").Parse()
	list2 := List("1 -> 2 -> 3").Parse()

	fmt.Println(list)
	fmt.Println(list2)
}

func Test_Slice(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"[2,1,0]", []int{2, 1, 0}},
		{"[2,1]", []int{2, 1}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.s, func(t *testing.T) {
			res := Slice(test.s).Ints()
			if !Equal(res, test.want) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}

func Test_ListEqual(t *testing.T) {
	tests := []struct {
		a *ListNode
		b *ListNode

		want bool
	}{
		{List("1->2->3").Parse(), List("1->2->3").Parse(), true},
		{List("1->2->3").Parse(), List("1->2").Parse(), false},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.a, test.b), func(t *testing.T) {
			res := listEqual(test.a, test.b)
			if res != test.want {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}

func Test_Equal(t *testing.T) {
	tests := []struct {
		a interface{}
		b interface{}

		want bool
	}{
		{List("1->2->3").Parse(), List("1->2->3").Parse(), true},
		{1, 1, true},
		{1, 2, false},
		{[]int{1, 2}, []int{1, 2}, true},
		{[2]int{1, 2}, [2]int{1, 2}, true},
		{[]int{1, 2}, []int{1}, false},
		{[2]int{1, 2}, [2]int{1}, false},
		{"a", "a", true},
		{"a", "b", false},
		{"a", 1, false},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.a, test.b), func(t *testing.T) {
			res := Equal(test.a, test.b)
			if res != test.want {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}

func Test_str2Tree(t *testing.T) {
	tests := []struct {
		str  string
		want *TreeNode
	}{
		{
			"[1,2,3,null,5,null,4]",
			nil,
		},
		{
			"[1,null,3,null]",
			nil,
		},
		{
			"[1,2,3,null, 5]",
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.str, func(t *testing.T) {
			res := Tree(test.str).Parse()
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
