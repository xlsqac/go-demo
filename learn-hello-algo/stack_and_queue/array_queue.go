//	Package stack_and_queue 基于数组的队列实现
//
// push pop peek O(1) toSlice O(n)
package stack_and_queue

type arrayQueue struct {
	num         []int // 存放数据的数组
	front       int   // 指向队列头部的指针
	queSize     int   // 队列的长度
	queCapacity int   // 队列的容量，也就是数组的长度
}

func newArrayQueue(capacity int) *arrayQueue {
	return &arrayQueue{
		num:         make([]int, capacity),
		front:       0,
		queSize:     0,
		queCapacity: capacity,
	}
}

func (queue *arrayQueue) size() int {
	return queue.queSize
}

func (queue *arrayQueue) isEmpty() bool {
	return queue.queSize == 0
}

func (queue *arrayQueue) push(value int) {
	if queue.queSize == queue.queCapacity {
		return
	}
	rear := (queue.front + queue.queSize) % queue.queCapacity
	queue.num[rear] = value
	queue.queSize++
}

func (queue *arrayQueue) pop() any {
	res := queue.peek()
	queue.queSize--
	queue.front = (queue.front + 1) % queue.queCapacity
	return res
}

func (queue *arrayQueue) peek() any {
	if queue.isEmpty() {
		return nil
	}
	return queue.num[queue.front]

}

func (queue *arrayQueue) toSlice() []int {
	res := make([]int, queue.queSize)
	for i := 0; i < queue.queSize; i++ {
		res[i] = queue.num[(queue.front+i)%queue.queCapacity]
	}
	return res
}
