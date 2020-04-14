package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 5)
	twitter.GetNewsFeed(1)

	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	twitter.GetNewsFeed(1)
	twitter.Unfollow(1, 2)
	twitter.GetNewsFeed(1)
}

func Test(t *testing.T) {
	var res []int

	tweets := [][]int{
		{1, 2, 7},
		{3, 5},
		{4, 6},
	}

	pos := make([]int, len(tweets))
	num := 0
	allDone := true
	for num < 10 {
		minPos := -1
		min := math.MaxInt32
		for i := 0; i < len(tweets); i++ {
			idx := pos[i]
			tweet := tweets[i]
			if pos[i] < len(tweets[i]) {
				if tweet[idx] < min {
					min = tweet[idx]
					minPos = i
				}
				allDone = false
			}
		}

		if minPos != -1 {
			pos[minPos]++
		}

		if allDone == true {
			break
		}

		res = append(res, min)
		num++

		allDone = true
	}
}

func Test_mergeK(t *testing.T) {
	list := mergeKLists2([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		//{Val: 2, Next: &ListNode{Val: 6}},
		nil,
	})

	fmt.Println(list)
}

func TestAddTwoNumbersII(t *testing.T) {
	//(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	node := addTwoNumbers(l1, l2)

	fmt.Println(node)
}

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		A   []int
		K   int
		out int
	}{
		{
			A:   []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			K:   2,
			out: 6,
		},
		{
			A:   []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			K:   3,
			out: 10,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.A), func(t *testing.T) {
			out := longestOnes(test.A, test.K)
			if test.out != out {
				t.Errorf("want %v, but got %v", test.out, out)
			}
		})
	}
}

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		N    int
		want string
	}{
		{1, "A"},
		{26, "Z"},
		{28, "AB"},
		{52, "AZ"},
		{701, "ZY"},
		{26 * 26, "YZ"},
		{26*26 + 1, "ZA"},
		{26 * 26 * 26, "YYZ"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.N), func(t *testing.T) {
			out := convertToTitle(test.N)
			if test.want != out {
				t.Errorf("want %v, but got %v", test.want, out)
			}
		})
	}
}
