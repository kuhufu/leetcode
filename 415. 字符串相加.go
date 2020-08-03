package leetcode

/*
* 415. 字符串相加
 */

func addStrings(num1 string, num2 string) string {
	idx1 := len(num1) - 1
	idx2 := len(num2) - 1

	ans := make([]byte, max(len(num1), len(num2))+1)
	ansIdx := len(ans) - 1

	carryBit := 0
	for idx1 >= 0 || idx2 >= 0 {
		v1 := 0
		if idx1 >= 0 {
			v1 = int(num1[idx1] - '0')
		}
		v2 := 0
		if idx2 >= 0 {
			v2 = int(num2[idx2] - '0')
		}

		ans[ansIdx] = byte((v1+v2+carryBit)%10) + '0'
		carryBit = (v1 + v2 + carryBit) / 10

		idx1--
		idx2--
		ansIdx--
	}

	if carryBit == 1 {
		ans[0] = '1'
		return string(ans)
	}

	return string(ans[1:])
}
