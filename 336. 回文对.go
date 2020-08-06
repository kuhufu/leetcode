package leetcode

/*
* 336. 回文对
* https://leetcode-cn.com/problems/palindrome-pairs/
 */

func check(s1, s2 string) bool {
	l := abs(len(s1) - len(s2))
	start := 0

	if len(s1) > len(s2) {
		start = len(s1) - l
	}

	lo := 0
	hi := len(s2) - 1
	for lo < len(s1) && hi >= 0 {
		if s1[lo] != s2[hi] {
			return false
		}
		lo++
		hi--
	}

	var isPalindrome1 func(string) bool
	isPalindrome1 = func(s string) bool {
		lo := 0
		hi := len(s) - 1
		for lo < hi {
			if s[lo] != s[hi] {
				return false
			}
			lo++
			hi--
		}

		return true
	}

	if len(s1) > len(s2) {
		return isPalindrome1(s1[start:])
	} else {
		return isPalindrome1(s2[start : start+l])
	}
}

func palindromePairs(words []string) [][]int {
	var ans [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if check(words[i], words[j]) {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

type Node struct {
	ch  [26]*Node
	end bool
}

func (n *Node) make(words []string) {
	for _, word := range words {
		n.insert(word)
	}
}

func (n *Node) insert(word string) {
	cur := n
	for _, ch := range word {
		if cur.ch[ch-'a'] == nil {
			cur.ch[ch-'a'] = &Node{}
		}
		cur = cur.ch[ch-'a']
	}

	cur.end = true
}

func (n *Node) findWord(word string) bool {
	cur := n
	for _, v := range word {
		if cur == nil {
			return false
		}
		cur = cur.ch[v-'a']
	}

	if cur != nil && cur.end {
		return true
	}

	return false
}
