// Package stack_and_queue
// 基于双向链表实现的双向队列，使用标准库 list 作为链表
// push O(1) pop O(1) peek O(1) toList O(n)
package stack_and_queue

import "container/list"

type LinkedListDeque struct {
	data *list.List
}

/* 构造函数 */
func newLinkedListDeque() *LinkedListDeque {
	return &LinkedListDeque{list.New()}
}

func (deque *LinkedListDeque) size() int {
	return deque.data.Len()
}

func (deque *LinkedListDeque) isEmpty() bool {
	return deque.data.Len() == 0
}

func (deque *LinkedListDeque) pushFirst(num int) {
	deque.data.PushFront(num)
}

func (deque *LinkedListDeque) pushLast(num int) {
	deque.data.PushBack(num)
}

func (deque *LinkedListDeque) peekElement(isFront bool) *list.Element {
	if deque.isEmpty() {
		return nil
	}
	element := deque.data.Back()
	if isFront {
		element = deque.data.Front()
	}
	return element
}

func (deque *LinkedListDeque) pop(isFront bool) any {
	element := deque.peekElement(isFront)
	num := element.Value
	deque.data.Remove(element)
	return num
}

func (deque *LinkedListDeque) popFirst() any {
	return deque.pop(true)
}

func (deque *LinkedListDeque) popLast() any {
	return deque.pop(false)
}

func (deque *LinkedListDeque) peekFirst() any {
	return deque.peekElement(true).Value

}

func (deque *LinkedListDeque) peekLast() any {
	return deque.peekElement(false).Value
}

func (deque *LinkedListDeque) toList() *list.List {
	return deque.data
}
