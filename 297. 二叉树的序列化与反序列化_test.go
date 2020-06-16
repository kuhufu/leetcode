package leetcode

import "testing"

func Test_serializeAndDeserialize(t *testing.T) {
	Run(t, func(s string) string {
		obj := TreeCodecConstructor()
		root := obj.deserialize(s)
		return obj.serialize(root)
	}, []Test{
		{
			Args{"[]"},
			Want{"[]"},
		},
		{
			Args{"[1,2,3,null,4]"},
			Want{"[1,2,3,null,4]"},
		},
	})
}
