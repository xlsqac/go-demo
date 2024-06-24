// Package stack_and_queue
// 基于数组实现的双向队列
// push pop peek O(1) toSlice O(n)
package stack_and_queue

type arrayDeque struct {
	nums          []int // 存放数据的数组
	front         int   // 指向队列头部的指针
	dequeSize     int   // 队列长度
	dequeCapacity int   // 队列容量，数组的长度
}

func newArrayDeque(capacity int) *arrayDeque {
	return &arrayDeque{
		nums:          make([]int, capacity),
		front:         0,
		dequeSize:     0,
		dequeCapacity: capacity,
	}
}

func (deque *arrayDeque) size() int {
	return deque.dequeSize
}

func (deque *arrayDeque) isEmpty() bool {
	return deque.dequeSize == 0
}

func (deque *arrayDeque) isFull() bool {
	return deque.dequeSize == deque.dequeCapacity
}

func (deque *arrayDeque) calIndex(index int) int {
	return (index + deque.dequeCapacity) % deque.dequeCapacity
}

func (deque *arrayDeque) push(num int, index int) {
	if deque.isFull() {
		panic("队列已满")
	}
	deque.dequeSize++
	deque.nums[index] = num
}

func (deque *arrayDeque) pushFirst(num int) {
	index := deque.calIndex(deque.front - 1)
	deque.push(num, index)
	deque.front = index
}

func (deque *arrayDeque) pushLast(num int) {
	index := deque.calIndex(deque.front + deque.dequeSize)
	deque.push(num, index)
}

func (deque *arrayDeque) popFirst() any {
	if deque.isEmpty() {
		panic("队列为空")
	}
	num := deque.nums[deque.front]
	deque.front = deque.calIndex(deque.front + 1)
	deque.dequeSize--
	return num
}

func (deque *arrayDeque) popLast() any {
	if deque.isEmpty() {
		panic("队列为空")
	}
	rear := deque.calIndex(deque.front + deque.dequeSize - 1)
	num := deque.nums[rear]
	deque.dequeSize--
	return num
}

func (deque *arrayDeque) peekFirst() any {
	if deque.isEmpty() {
		panic("队列为空")
	}
	num := deque.nums[deque.front]
	return num
}

func (deque *arrayDeque) peekLast() any {
	if deque.isEmpty() {
		panic("队列为空")
	}
	rear := deque.calIndex(deque.front + deque.dequeSize - 1)
	num := deque.nums[rear]
	return num
}

func (deque *arrayDeque) toSlice() []int {
	res := make([]int, deque.dequeSize)
	for i := 0; i < deque.dequeSize; i++ {
		index := deque.calIndex(deque.front + i)
		res[i] = deque.nums[index]
	}
	return res

}
