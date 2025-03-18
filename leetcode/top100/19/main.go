// package main
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/?envType=study-plan-v2&envId=top-100-liked
// 删除链表的倒数第 N 个节点
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	{
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}
		node3 := &ListNode{Val: 3}
		node4 := &ListNode{Val: 4}
		node5 := &ListNode{Val: 5}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5

		head := removeNthFromEndDoublePoint(node1, 2)
		printListNode(head)
		fmt.Println()
	}
	{
		node1 := &ListNode{Val: 1}
		head := removeNthFromEndDoublePoint(node1, 1)
		printListNode(head)
		fmt.Println()
	}
	{
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}

		node1.Next = node2

		head := removeNthFromEndDoublePoint(node1, 1)
		printListNode(head)
		fmt.Println()
	}
	{
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}

		node1.Next = node2

		head := removeNthFromEndDoublePoint(node1, 2)
		// fmt.Print(node1, node2, head)
		printListNode(head)
	}
}

// removeNthFromEnd 删除链表的倒数第 N 个节点
// 时间复杂度 O(n)，空间复杂度 O(n)
// 利用切片把链表的值存储起来，然后根据索引删除对应的节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	length := len(nodes)
	deleteIndex := length - n
	// 只有一个元素的情况
	if deleteIndex == 0 && length == 1 {
		return nil
	}
	// 要删除最后一个元素
	if n == 1 {
		nodes[deleteIndex-1].Next = nil
		return head
	} else if n == length {
		// 要删除第一个元素
		return nodes[1]

	}
	nodes[deleteIndex-1].Next = nodes[deleteIndex+1]
	return head
}

// removeNthFromEndDoublePoint 删除链表的倒数第 N 个节点，双指针法
// 时间复杂度 O(n)，空间复杂度 O(1)
// 两个指针，一个指针先走 n 步，然后两个指针一起走，当第一个指针走到最后一个节点时，第二个指针就是要删除的节点
func removeNthFromEndDoublePoint(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// printListNode 打印链表中的值
func printListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
