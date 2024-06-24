// Package stack_and_queue  基于链表的栈，利用标准库 list 实现
// Push，Pop, Peek O(1), ToList O(n)
package stack_and_queue

import "container/list"

type LinkedListStack struct {
	data *list.List
}

// NewLinkedListStack /* 初始化 */
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		data: list.New(),
	}
}

// Push  入栈
func (s *LinkedListStack) Push(value int) {
	s.data.PushBack(value)
}

// Pop  出栈
func (s *LinkedListStack) Pop() any {
	e := s.peekElement()
	s.data.Remove(e)
	return e.Value
}

// Peek  获取栈顶
func (s *LinkedListStack) Peek() any {
	return s.peekElement().Value
}

// peekElement 获取栈顶的结点
func (s *LinkedListStack) peekElement() *list.Element {
	if s.IsEmpty() {
		return nil
	}
	e := s.data.Back()
	return e
}

func (s *LinkedListStack) Size() int {
	return s.data.Len()
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.data.Len() == 0
}

func (s *LinkedListStack) ToList() *list.List {
	return s.data
}
