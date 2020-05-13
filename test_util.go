package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type Args = []interface{}
type Want = []interface{}

type Test struct {
	Args Args
	Want Want
}

func Run(t *testing.T, fn0 interface{}, tests []Test) {
	fn := reflect.ValueOf(fn0)
	if fn.Kind() != reflect.Func {
		panic(fmt.Sprintf("Input.Fn:%v is not Func type", fn.Type()))
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.Args...), func(t *testing.T) {
			res := toInterfaceSlice(fn.Call(buildArgs(test.Args)))
			if !Equal(test.Want, res) {
				t.Errorf("want %v, but got %v", test.Want, res)
			}
		})
	}
}

func buildArgs(args []interface{}) []reflect.Value {
	var res []reflect.Value

	for _, arg := range args {
		res = append(res, reflect.ValueOf(arg))
	}

	return res
}

func toInterfaceSlice(a []reflect.Value) []interface{} {
	var res []interface{}

	for _, v := range a {
		res = append(res, v.Interface())
	}

	return res
}
