// Package main 堆
// 完全二叉树，用数组实现
// 出堆、入堆 O(logn)
// peek, size, isEmpty O(1)
// 建堆、堆化
package main

import "fmt"

type MaxHeap struct {
	data []any
}

// Left 获取当前索引的节点的左子节点的索引
func (h *MaxHeap) Left(i int) int {
	return 2*i + 1
}

// Right 获取右子节点的索引
func (h *MaxHeap) Right(i int) int {
	return 2*i + 2
}

// 获取父节点的索引
func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

// peek 访问堆顶元素 O(1)
func (h *MaxHeap) peek() any {
	return h.data[0]
}

// swap 交换节点
func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]

}

// push 入堆 O(logn)
func (h *MaxHeap) push(val any) {
	h.data = append(h.data, val)
	// 堆化
	h.siftUp(h.maxIndex())
}

// siftUp 从底到顶堆化，插入元素后，成立堆的条件可能被破坏，修复插入节点到根节点的路径上的各个节点
func (h *MaxHeap) siftUp(i int) {
	for {
		parent := h.parent(i)
		// 当越过根节点或节点无需修复时退出堆化
		if parent < 0 || h.data[i].(int) <= h.data[parent].(int) {
			break
		}
		// 交换两个节点
		h.swap(i, parent)
		i = parent
	}

}

// pop 堆顶出堆 logn
// 1. 交换堆顶元素和堆底元素
// 2. 删除堆底元素(也就是堆顶元素)
// 3. 从根节点开始堆化
func (h *MaxHeap) pop() any {
	if h.isEmpty() {
		fmt.Println("堆为空")
		return nil
	}
	// 堆顶索引
	peekIndex := 0
	// 堆底索引
	tailIndex := h.maxIndex()

	peek := h.data[peekIndex]
	// 交换堆顶和堆底
	h.swap(peekIndex, tailIndex)
	// 删除堆底元素
	h.data = h.data[:tailIndex]
	// 从顶到底堆化
	h.siftDown(peekIndex)

	return peek
}

// siftDown 从顶到底堆化
func (h *MaxHeap) siftDown(i int) {
	for {
		left, right, maxNode, maxValue := h.Left(i), h.Right(i), i, h.data[i]
		if i < h.size() && h.data[left].(int) > maxValue.(int) {
			maxNode = left
		}
		if i < h.size() && h.data[right].(int) > maxValue.(int) {
			maxNode = right
		}
		// 索引越界或者最大就是 i 时，说明堆化完成了
		if maxNode == i {
			break
		}
		h.swap(i, maxNode)
		i = maxNode
	}

}

// size 获取堆元素数量 O(1)
func (h *MaxHeap) size() int {
	return len(h.data)
}

// isEmpty 判断堆是否为空堆
func (h *MaxHeap) isEmpty() bool {
	return h.size() == 0
}

// MaxIndex 获取数组的最大索引
func (h *MaxHeap) maxIndex() int {
	return h.size() - 1
}
