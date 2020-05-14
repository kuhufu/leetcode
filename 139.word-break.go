package leetcode

import "strings"

/*
* 139. 单词拆分
* https://leetcode-cn.com/problems/word-break/
 */

/*
* 难度：中等
 */

func wordBreak(s string, wordDict []string) bool {
	mem := make([]bool, len(s)+1)
	mem[0] = true
	for i := 0; i < len(s)+1; i++ {
		str := s[:i]
		for _, word := range wordDict {
			if strings.HasSuffix(str, word) && mem[i-len(word)] {
				mem[i] = true
			}
		}
	}

	return mem[len(s)]
}
