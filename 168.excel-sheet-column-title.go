package leetcode

/*
* 168. Excel表列名称
* https://leetcode-cn.com/problems/excel-sheet-column-title/solution/
 */

//简单题里的难题，还不是一般的难
//看leetcode里的题解 https://leetcode-cn.com/problems/excel-sheet-column-title/solution/xiang-xi-tong-su-de-si-lu-fen-xi-by-windliang-2/
func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		res = string((n-1)%26+'A') + res
		n = (n - 1) / 26
	}
	return res
}
