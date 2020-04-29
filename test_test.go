package leetcode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	bytes, err := json.Marshal(2)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
