// Package main 堆
// 完全二叉树，用数组实现
// 大顶堆和小顶堆，堆顶和堆底
// 出堆、入堆 O(logn)
// peek, size, isEmpty O(1)
// 建堆、堆化
// 使用 heap.Interface 实现大顶堆
package main

import (
	"container/heap"
	"fmt"
)

type intHeap []any

// Push heap.Interface 的方法
func (h *intHeap) Push(x any) {
	*h = append(*h, x)
}

// Pop heap.Interface
func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface
func (h *intHeap) Len() int {
	return len(*h)
}

// Less sort.Interface
func (h *intHeap) Less(i, j int) bool {
	return (*h)[i].(int) > (*h)[j].(int)
}

// Swap sort.Interface
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}

func main() {
	/* 初始化堆 */
	// 初始化大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	/* 获取堆顶元素 */
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为 %d\n", top)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	heap.Pop(maxHeap) // 5
	heap.Pop(maxHeap) // 4
	heap.Pop(maxHeap) // 3
	heap.Pop(maxHeap) // 2
	heap.Pop(maxHeap) // 1

	/* 获取堆大小 */
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	/* 判断堆是否为空 */
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
