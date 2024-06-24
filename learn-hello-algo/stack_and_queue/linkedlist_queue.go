// Package stack_and_queue 基于链表的队列实现，使用标准库 list 作为链表
// push O(1) pop O(1) peek O(1) toList O(n)
package stack_and_queue

import "container/list"

type LinkedListQueue struct {
	data *list.List
}

/* 构造函数 */
func newLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		data: list.New(),
	}
}

/* 队列长度 */
func (queue *LinkedListQueue) size() int {
	return queue.data.Len()
}

/* 队列是否为空 */
func (queue *LinkedListQueue) isEmpty() bool {
	return queue.size() == 0
}

/* 入队，从队尾增加元素 */
func (queue *LinkedListQueue) push(value int) {
	queue.data.PushBack(value)
}

/* 出队，从队首删除元素 */
func (queue *LinkedListQueue) pop() any {
	e := queue.peekNode()
	queue.data.Remove(e)
	return e.Value
}

/* 获取队首元素 */
func (queue *LinkedListQueue) peek() any {
	e := queue.peekNode()
	return e.Value
}

/* 获取队首结点 */
func (queue *LinkedListQueue) peekNode() *list.Element {
	if queue.isEmpty() {
		return nil
	}
	e := queue.data.Front()
	return e
}

/* 获取 list 用于打印 */
func (queue *LinkedListQueue) toList() *list.List {
	return queue.data
}
