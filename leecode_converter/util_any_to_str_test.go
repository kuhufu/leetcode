package leecode_converter

import (
	"fmt"
	"testing"
)

func Test_graph2Str(t *testing.T) {
	tests := []struct {
		Graph *Node
		want  string
	}{
		{Graph("[[2,4],[1,3],[2,4],[1,3]]"), "[[2,4],[1,3],[2,4],[1,3]]"},
		{Graph("[[2,4],[1,3],[2,4],[1,3]]"), "[[2,4],[1,3],[2,4],[1,3]]"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.want, func(t *testing.T) {
			res := graph2Str(test.Graph)
			if res != test.want {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}

func Test_list2Str(t *testing.T) {
	tests := []struct {
		list *ListNode
		want string
	}{
		{List("1->2->3"), "1->2->3"},
		{List("1 -> 2 -> 3"), "1->2->3"},
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

func Test_listArr2str(t *testing.T) {
	tests := []struct {
		listArr []*ListNode
		want    string
	}{
		{
			[]*ListNode{
				List("1->4->5"),
				List("1->3->4"),
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
