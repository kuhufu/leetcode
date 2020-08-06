package leecode_converter

import (
	"encoding/json"
	"fmt"
)

type SliceConverter string

func (s SliceConverter) Ints() []int {
	var res []int

	if err := json.Unmarshal([]byte(s), &res); err != nil {
		fmt.Println(s)
		panic(err)
	}

	return res
}

func (s SliceConverter) DInts() [][]int {
	var res [][]int

	if err := json.Unmarshal([]byte(s), &res); err != nil {
		fmt.Println(s)
		panic(err)
	}

	return res
}

func (s SliceConverter) Strings() []string {
	var res []string

	if err := json.Unmarshal([]byte(s), &res); err != nil {
		fmt.Println(s)
		panic(err)
	}

	return res
}

func (s SliceConverter) DStrings() [][]string {
	var res [][]string

	if err := json.Unmarshal([]byte(s), &res); err != nil {
		fmt.Println(s)
		panic(err)
	}

	return res
}
