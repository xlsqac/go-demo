package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func newListNode(value int) *ListNode {
	return &ListNode{
		Value: value,
		Next:  nil,
	}

}

func traverse(n0 *ListNode) (res []int) {
	// 遍历链表
	for n0 != nil {
		res = append(res, n0.Value)
		n0 = n0.Next
	}
	return res
}

func insertNode(n0 *ListNode, p *ListNode) {
	// 插入结点 O(1)
	p.Next = n0.Next
	n0.Next = p
}

func deleteNode(n0 *ListNode) {
	// 删除结点 O(1)
	if n0.Next == nil {
		return
	}
	p := n0.Next
	n0.Next = p.Next

}

func accessNode(head *ListNode, index int) *ListNode {
	// 访问结点 O(n)
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

func findNode(head *ListNode, target int) int {
	// 查找结点 O(n)
	index := 0
	for head != nil {
		if head.Value == target {
			return index
		}
		index += 1
		head = head.Next
	}
	return -1
}

func main() {
	// 初始化列表
	n0 := newListNode(1)
	n1 := newListNode(3)
	n2 := newListNode(2)
	n3 := newListNode(5)
	n4 := newListNode(4)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	// [1,3,2,5,4]
	fmt.Printf("[init linekd list] %v\n", traverse(n0))

	// 插入
	p := newListNode(0)
	// [1,0,3,2,5,4]
	insertNode(n0, p)
	fmt.Printf("[insert node] %v\n", traverse(n0))

	// 删除
	// [1,0,3,2,4]
	deleteNode(n2)
	fmt.Printf("[delete node] %v\n", traverse(n0))

	// 访问
	// 3
	p = accessNode(n0, 2)
	fmt.Printf("[access node index 2] %d\n", p.Value)

	// 查找
	// 3
	index := findNode(n0, 2)
	fmt.Printf("[find node value 2] %d\n", index)

}
