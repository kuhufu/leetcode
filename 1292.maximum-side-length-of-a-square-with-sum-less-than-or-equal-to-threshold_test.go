package leetcode

import (
	"testing"
)

func Test_maxSideLen(t *testing.T) {
	tests := []struct {
		matrix    [][]int
		threshold int
		want      int
	}{
		{
			Str22DSlice("[[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]]"),
			4,
			2,
		},
		{
			Str22DSlice("[[21,91,65,94,46,95,70,26,88,29,73,24,62,32,35,49,79,67,23,45,6,7,60,98,31,69,86,64,28,78,47,57,20,17],[67,74,78,10,19,22,17,58,44,31,100,37,30,16,6,56,81,60,97,41,63,87,40,79,26,77,86,80,15,39,23,28,34,83],[24,73,55,82,0,62,89,54,49,28,77,43,81,12,84,91,72,25,5,29,68,23,93,42,34,13,69,94,35,19,37,36,78,74],[81,27,83,11,31,0,85,71,35,4,8,92,62,99,61,47,63,17,93,7,43,52,60,33,79,41,40,76,13,80,57,39,51,54],[21,81,28,100,47,83,5,62,60,97,86,90,85,46,0,42,72,58,74,77,38,69,41,35,95,54,8,99,1,87,33,51,66,27],[64,69,54,34,22,21,83,82,85,2,5,81,67,56,16,4,77,84,40,72,70,95,51,89,96,61,63,7,1,28,99,62,75,47],[2,15,95,54,89,10,96,27,6,85,1,25,81,33,37,18,39,63,49,72,77,69,48,31,24,22,87,66,8,53,17,74,86,68],[68,88,5,53,94,65,38,95,48,37,42,55,29,93,80,47,36,77,64,46,43,2,99,89,72,61,9,16,74,20,76,73,82,11],[28,57,76,12,86,31,37,62,70,96,27,17,13,77,59,41,85,55,94,97,3,34,56,67,74,47,33,38,6,64,51,63,83,80],[66,41,13,1,70,81,15,0,82,92,22,12,44,74,16,34,26,20,36,35,77,53,52,79,56,29,59,10,3,89,48,80,55,24],[92,15,98,32,23,62,68,93,1,17,53,28,21,54,60,49,51,43,52,77,3,34,46,10,84,35,75,89,26,24,80,45,31,22],[83,95,2,51,17,7,28,89,5,80,52,43,67,91,57,86,40,53,90,11,65,38,68,47,33,82,63,85,61,37,15,22,87,19],[2,99,69,61,0,20,50,77,48,11,22,17,28,66,1,70,98,74,88,82,63,34,75,76,78,24,23,8,10,35,26,85,19,57],[94,18,44,87,91,81,3,15,31,40,4,14,63,47,89,49,12,16,9,23,50,66,0,28,43,10,39,20,78,56,73,99,95,5],[94,26,89,14,69,27,8,85,87,12,24,35,44,81,15,56,38,63,75,60,99,47,32,76,10,19,62,93,20,98,70,21,3,30],[53,35,93,3,32,84,85,68,64,77,86,59,2,65,47,12,76,5,42,61,7,37,50,0,38,79,18,52,34,82,29,55,41,21],[96,25,50,55,69,19,15,9,30,88,16,32,38,41,80,6,97,79,28,27,65,47,89,0,73,84,3,51,11,33,1,72,23,92],[30,65,75,51,17,25,60,100,36,66,83,16,5,26,96,37,99,93,10,52,88,19,2,87,55,47,7,28,91,21,35,85,62,31],[63,96,6,57,68,71,60,81,64,4,70,72,1,74,34,80,17,67,28,54,79,32,23,12,73,39,89,93,55,100,40,24,58,45],[93,41,50,10,73,87,21,83,77,4,85,1,3,24,29,30,55,2,48,31,11,38,34,44,96,62,49,68,15,46,12,56,59,81],[48,47,87,10,30,78,23,44,11,58,39,7,9,53,20,46,73,77,40,69,2,62,21,88,75,43,90,14,36,13,85,91,54,28],[64,92,98,7,94,70,81,80,89,38,73,85,39,95,57,65,69,60,32,97,75,77,30,90,3,37,2,53,43,72,36,61,100,58],[46,47,90,10,12,50,83,98,75,25,55,11,51,0,60,86,63,28,79,91,72,78,88,14,67,92,9,19,54,85,57,99,7,38],[13,14,28,69,34,1,97,23,53,91,9,75,58,61,21,24,50,95,79,93,52,26,72,86,89,40,96,65,90,45,15,87,18,63],[89,100,79,64,55,40,32,65,57,93,23,49,29,54,56,2,14,34,44,73,52,47,25,33,20,51,76,81,18,21,98,96,7,97],[54,64,75,50,35,32,10,68,51,4,65,79,28,95,98,57,41,26,7,18,19,6,93,63,0,40,91,27,1,49,30,37,12,73],[83,73,24,70,8,72,6,21,80,93,43,13,82,33,99,37,26,40,69,64,11,2,81,53,46,38,55,76,62,74,92,49,30,4],[63,59,70,95,98,86,72,67,71,37,56,2,29,50,93,39,44,92,15,34,22,40,7,84,43,35,96,54,53,90,45,81,73,38],[32,79,50,29,83,16,54,90,66,19,51,96,92,87,9,93,94,76,23,56,12,95,98,31,15,97,30,71,42,81,7,43,52,88],[65,53,23,63,82,81,9,67,75,44,39,69,16,55,99,93,17,71,36,96,84,85,28,19,35,79,70,1,25,26,72,80,22,10],[3,75,58,74,18,46,16,88,21,10,1,29,76,93,73,98,40,32,52,62,22,26,14,66,38,57,55,34,82,72,60,83,90,19],[74,48,36,32,90,80,98,17,46,75,10,30,2,19,9,44,33,5,14,61,79,89,96,23,71,81,100,7,56,72,11,18,63,84],[93,81,35,27,18,95,13,44,23,49,54,86,87,96,19,8,57,24,62,66,74,3,32,78,46,17,22,53,30,56,25,6,61,9],[6,41,29,24,2,19,77,71,57,63,81,36,73,12,49,22,4,95,61,47,21,80,53,14,9,42,51,40,91,100,65,48,97,27],[75,4,70,8,58,25,50,2,30,93,22,24,29,31,87,57,39,71,73,69,94,81,48,100,9,14,60,27,54,36,77,47,92,68],[15,48,8,83,28,41,61,52,3,23,0,67,55,87,60,73,46,35,47,25,71,95,69,31,54,1,76,10,94,13,58,14,51,93],[27,54,43,13,68,76,82,88,19,47,8,100,36,58,72,39,84,63,16,31,33,57,60,30,99,69,98,10,24,38,44,86,29,67],[19,75,51,31,16,32,90,27,0,15,69,80,21,59,76,46,60,30,79,2,33,81,44,37,77,89,45,1,61,28,41,50,3,47],[72,19,99,80,33,86,2,96,100,22,85,74,55,18,42,70,21,73,76,95,90,26,51,75,38,41,3,69,49,31,7,61,68,53],[46,55,94,14,25,16,22,87,73,97,43,62,57,51,37,28,86,95,3,1,31,59,40,7,4,24,34,10,56,61,96,72,69,99],[26,3,36,72,62,83,52,60,99,27,57,11,86,64,10,70,43,73,29,48,67,68,87,20,92,14,25,47,8,17,22,21,49,78],[12,3,10,76,97,87,57,30,42,20,9,73,27,23,51,35,92,19,46,68,70,24,61,85,11,77,72,59,5,95,80,67,50,52],[25,28,17,10,76,56,99,96,54,83,15,12,71,48,84,49,58,14,43,21,88,75,68,98,29,85,73,22,89,50,55,72,30,5],[72,83,17,79,21,0,68,73,91,8,78,75,52,13,11,60,53,87,26,80,85,38,16,89,22,84,34,41,47,39,37,98,54,19],[42,39,16,8,44,36,96,63,79,83,5,33,98,7,64,35,49,10,41,40,18,73,65,70,23,47,55,50,28,4,26,3,100,13],[73,27,87,76,64,13,39,32,71,23,42,60,59,94,99,47,44,81,5,82,97,89,80,86,41,10,91,4,54,67,63,98,22,7],[39,47,96,35,29,72,18,19,95,76,85,71,63,88,36,16,10,55,45,62,61,69,99,65,50,81,41,7,14,59,33,97,23,98],[56,38,52,81,50,42,57,26,82,72,58,10,77,25,73,13,69,6,7,30,74,70,0,34,41,29,84,17,1,55,66,9,60,44],[31,96,45,52,79,33,95,55,32,98,89,26,69,24,87,58,29,6,30,40,23,70,62,15,9,17,0,21,91,48,88,12,3,11],[62,48,8,64,44,81,84,6,47,96,91,60,24,57,65,80,11,68,54,46,71,19,82,74,26,99,53,100,94,98,38,56,10,85],[77,97,12,72,92,73,6,85,9,95,11,83,54,51,61,66,20,16,39,90,93,5,8,47,75,32,22,35,33,60,56,18,38,0],[75,25,21,9,53,55,26,59,62,84,88,46,3,34,41,65,64,10,6,61,93,40,4,91,27,24,45,35,2,0,72,97,12,31],[98,24,92,16,29,17,42,75,76,7,88,66,80,3,84,72,5,68,67,45,18,41,32,46,78,23,40,60,21,63,9,53,77,10]]"),
			51510,
			32,
		},
		{
			Str22DSlice("[[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]]"),
			1,
			0,
		},
		{
			Str22DSlice("[[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]]"),
			40184,
			2,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(Str(test.matrix), func(t *testing.T) {
			res := maxSideLength(test.matrix, test.threshold)
			if !Equal(test.want, res) {
				t.Errorf("want %v, but got %v", test.want, res)
			}
		})
	}
}
