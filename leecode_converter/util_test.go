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

func Test_List(t *testing.T) {

	list := List("1->2->3").Parse()
	list2 := List("1 -> 2 -> 3").Parse()

	fmt.Println(list)
	fmt.Println(list2)
}

func Test_list2Str(t *testing.T) {
	tests := []struct {
		list *ListNode
		want string
	}{
		{List("1->2->3").Parse(), "1->2->3"},
		{List("1 -> 2 -> 3").Parse(), "1->2->3"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.want, func(t *testing.T) {
			res := list2Str(test.list)
			if res != test.want {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
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

func Test_slice2Str(t *testing.T) {
	tests := []struct {
		slice []int
		want  string
	}{
		{[]int{2, 1, 0}, "[2,1,0]"},
		{[]int{2, 1}, "[2,1]"},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.slice), func(t *testing.T) {
			res := slice2Str(test.slice)
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

func Test_listArr2str(t *testing.T) {
	tests := []struct {
		listArr []*ListNode
		want    string
	}{
		{
			[]*ListNode{
				List("1->4->5").Parse(),
				List("1->3->4").Parse(),
			},
			"[1->4->5,1->3->4]",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(listArr2Str(test.listArr), func(t *testing.T) {
			res := listArr2Str(test.listArr)
			if !Equal(test.want, res) {
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

func Test_Str(t *testing.T) {
	tests := []struct {
		args []interface{}
		want string
	}{
		{
			[]interface{}{1, 2, 3},
			"1, 2, 3",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.args), func(t *testing.T) {
			res := Str(test.args...)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
