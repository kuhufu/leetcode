package leetcode

/*
* 202. 快乐数
* https://leetcode-cn.com/problems/happy-number/
 */

/*
* 难度: 简单
 */

//方法一：用 HashSet 检测循环
func isHappy(n int) bool {
	s := map[int]struct{}{
		n: {},
	}

	getNext := func(n int) int {
		sum := 0
		for ; n != 0; n /= 10 {
			mod := n % 10
			sum += mod * mod
		}
		return sum
	}

	for n != 1 {
		n = getNext(n)
		if _, ok := s[n]; ok {
			return false
		}
		s[n] = struct{}{}
	}

	return true
}

//方法二：快慢指针法
func isHappy2(n int) bool {
	getNext := func(n int) int {
		sum := 0
		for ; n != 0; n /= 10 {
			mod := n % 10
			sum += mod * mod
		}
		return sum
	}

	num := n
	fasterNum := getNext(num)

	for num != 1 && num != fasterNum {
		num = getNext(num)
		fasterNum = getNext(getNext(fasterNum))
	}

	return num == 1
}
