package coding_interviews_2th

/*
* 面试题56 - I. 数组中数字出现的次数
* https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
 */

/*
* 难度: 中等
* 通过: true
 */

//这题用位运算来分组是真的妙啊，妙蛙种子的妙
func singleNumbers(nums []int) []int {
	v := 0
	for _, num := range nums {
		v ^= num
	}

	mask := 1

	for v&mask == 0 {
		mask <<= 1
	}

	a := 0
	b := 0

	for _, num := range nums {
		if num&mask == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
