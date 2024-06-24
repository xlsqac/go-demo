// Package stack_and_queue  基于数组的栈，利用 slice 实现
package stack_and_queue

type ArrayStack struct {
	data []int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: []int{},
	}
}

func (s *ArrayStack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *ArrayStack) Pop() any {
	v := s.Peek()
	s.data = s.data[:s.Size()-1]
	return v
}

func (s *ArrayStack) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.Size()-1]
}

func (s *ArrayStack) Size() int {
	return len(s.data)
}

func (s *ArrayStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *ArrayStack) ToSlice() []int {
	return s.data
}
