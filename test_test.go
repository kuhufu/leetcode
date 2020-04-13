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
