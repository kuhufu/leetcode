package cracking_the_coding_interview_6th

import (
	"fmt"
	"math"
)

/*
* 面试题 17.13. 恢复空格
* https://leetcode-cn.com/problems/re-space-lcci/
 */

func respace(dictionary []string, sentence string) int {
	//return respace1(dictionary, sentence)
	return respace2(dictionary, sentence)
}

// 方法一：Trie + 动态规划
func respace1(dictionary []string, sentence string) int {
	n, inf := len(sentence), 0x3f3f3f3f
	root := &Trie{next: [26]*Trie{}}
	for _, word := range dictionary {
		root.insert(word)
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		curPos := root
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1] - 'a')
			if curPos.next[t] == nil {
				break
			} else if curPos.next[t].isEnd {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			curPos = curPos.next[t]
		}
	}
	return dp[n]
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (t *Trie) insert(s string) {
	curPos := t
	for i := len(s) - 1; i >= 0; i-- {
		t := int(s[i] - 'a')
		if curPos.next[t] == nil {
			curPos.next[t] = &Trie{next: [26]*Trie{}}
		}
		curPos = curPos.next[t]
	}
	curPos.isEnd = true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

const (
	P    = math.MaxInt32
	BASE = 41
)

// 方法二：字符串哈希
func respace2(dictionary []string, sentence string) int {
	hashValues := map[int]string{}
	for _, word := range dictionary {
		if v := hashValues[getHash(word)]; v != "" {
			fmt.Println(v, word)
		}
		hashValues[getHash(word)] = word
	}

	dp := make([]int, len(sentence)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = len(sentence)
	}
	for i := 1; i <= len(sentence); i++ {
		dp[i] = dp[i-1] + 1
		hashValue := 0
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1]-'a') + 1 // +1 是必须的
			hashValue = (hashValue*BASE + t) % P
			if word, ok := hashValues[hashValue]; ok {
				s := sentence[j-1 : i]
				if word == s {
					dp[i] = min(dp[i], dp[j-1])
				}
				if word != s {
					fmt.Println(word, s)
				}
			}
		}
	}
	return dp[len(sentence)]
}

func getHash(s string) int {
	hashValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		hashValue = (hashValue*BASE + int(s[i]-'a') + 1) % P // +1 是必须的，不然 a, aa 值相同，b, ba, baa, baaa 值相同
	}
	return hashValue
}
