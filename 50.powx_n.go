package leetcode

/*
* 50. Pow(x, n)
* https://leetcode-cn.com/problems/powx-n/
 */

/*
* 难度: 中等
 */

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return x * half * half
	}
}
