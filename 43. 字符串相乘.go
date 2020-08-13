package leetcode

import (
	"reflect"
	"strings"
	"unsafe"
)

/*
* 43. 字符串相乘
* https://leetcode-cn.com/problems/multiply-strings/
 */

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := "0"
	for i := 0; i < len(num2); i++ {
		s := num2[len(num2)-1-i]
		v := multi(num1, s) + generateZero(i)
		ans = add(ans, v)
	}

	return ans
}

func multi(num string, c byte) string {
	n := int(c - '0')
	if n == 0 {
		return "0"
	}

	if n == 1 {
		return num
	}

	carryBit := 0

	ans := make([]byte, len(num)+1)
	idx := len(ans) - 1

	for i := len(num) - 1; i >= 0; i-- {
		v := int(num[i]-'0')*n + carryBit
		carryBit = v / 10
		v = v % 10

		ans[idx] = byte(v + '0')
		idx--
	}

	if carryBit != 0 {
		ans[idx] = byte(carryBit + '0')
		return string(ans)
	}

	return string(ans[1:])
}

func add(num1 string, num2 string) string {

	idx1 := len(num1) - 1
	idx2 := len(num2) - 1

	ans := make([]byte, max(len(num1), len(num2))+1)
	ansIdx := len(ans) - 1

	carryBit := 0
	for idx1 >= 0 || idx2 >= 0 {
		v1, v2 := 0, 0
		if idx1 >= 0 {
			v1 = int(num1[idx1] - '0')
		}
		if idx2 >= 0 {
			v2 = int(num2[idx2] - '0')
		}

		v := v1 + v2 + carryBit
		carryBit = v / 10
		v = v % 10

		ans[ansIdx] = byte(v + '0')

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

func generateZero(n int) string {
	builder := strings.Builder{}
	builder.Grow(n)

	for i := 0; i < n; i++ {
		builder.WriteByte('0')
	}

	return builder.String()
}

func bytes2String(bytes []byte) string {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

	s := reflect.StringHeader{
		Data: p.Data,
		Len:  p.Len,
	}

	return *(*string)(unsafe.Pointer(&s))
}
