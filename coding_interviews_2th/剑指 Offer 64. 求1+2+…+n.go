package coding_interviews_2th

/*
* 面试题64. 求1+2+…+n
* https://leetcode-cn.com/problems/qiu-12n-lcof/
 */

func sumNums(n int) (ans int) {
	arr := make([]struct{}, n)

	defer func() {
		_ = recover()
	}()

	var f func(arr []struct{})
	f = func(arr []struct{}) {
		ans += len(arr)
		f(arr[:len(arr)-1])
	}

	f(arr)
	return
}

func sumNums2(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
