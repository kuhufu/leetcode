package leetcode

/*
* 69. x 的平方根
* https://leetcode-cn.com/problems/sqrtx/
 */

func mySqrt(x int) int {
	lo := 0
	hi := x

	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid > x {
			hi = mid - 1
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo - 1
}
