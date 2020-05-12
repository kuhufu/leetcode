package leetcode

/*
* 155. 最小栈
* https://leetcode-cn.com/problems/min-stack/
 */

//这题的解法挺有意思的
type MinStack struct {
	data []int
	min  []int
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)

	if len(s.min) != 0 && s.min[len(s.min)-1] < x {
		x = s.min[len(s.min)-1]
	}
	s.min = append(s.min, x)
}

func (s *MinStack) Pop() {
	if len(s.data) == 0 {
		return
	}

	s.data = s.data[:len(s.data)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.min[len(s.min)-1]
}
