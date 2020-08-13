package leetcode

import (
	"reflect"
	"unsafe"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func IsDigital(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsAlphabet(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func bytes2String(bytes []byte) string {
	return *(*string)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes))))
}

func string2Bytes(s string) []byte {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))

	cp := reflect.SliceHeader{
		Data: p.Data,
		Len:  p.Len,
		Cap:  p.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&cp))
}
