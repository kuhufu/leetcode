package leetcode

/*
* 796. 旋转字符串
* https://leetcode-cn.com/problems/rotate-string/
 */

/*
* 难度：简单
 */

//这题还是挺有意思的
func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(B) == 0 && len(A) == 0 {
		return true
	}

	t := B[0]
	for i, c := range []byte(A) {
		if c == t {
			if A[i:] == B[:len(A)-i] && A[:i] == B[len(A)-i:] {
				return true
			}
		}
	}

	return false
}
