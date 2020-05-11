package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			2,
			5,
			32,
		},
	}

	for _, test := range tests {
		t.Run(Str(test.x, test.n), func(t *testing.T) {
			res := myPow(test.x, test.n)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
