package leetcode

/*
* 17. 电话号码的字母组合
* https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
 */

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var ans []string

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var f func([]byte, string)
	f = func(v []byte, digits string) {
		if digits == "" {
			ans = append(ans, string(v))
			return
		}

		for _, c := range m[digits[0]] {
			f(append(v, byte(c)), digits[1:])
		}
	}

	f(make([]byte, 0, len(digits)), digits)
	return ans
}
