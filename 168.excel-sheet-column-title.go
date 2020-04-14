package leetcode

/*
* 168. Excel表列名称
* https://leetcode-cn.com/problems/excel-sheet-column-title/solution/
*/

//简单题里的难题，还不是一般的难
//看leetcode里的题解 https://leetcode-cn.com/problems/excel-sheet-column-title/solution/xiang-xi-tong-su-de-si-lu-fen-xi-by-windliang-2/
func convertToTitle(n int) string {
	ans := ""

	for n > 0 {
		c := n % 26
		if c == 0 {
			c = 26
			n -= 1
		}
		ans = string('A'+c-1) + ans

		n /= 26
	}

	return ans
}
