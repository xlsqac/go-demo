// package main
// https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan-v2&envId=top-100-liked
// 两两交换链表中的节点
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

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4

		head := swapPairsRecursion(node1)
		printLinkedList(head)
	}
}

// swapPairsRecursion 两两交换连表中的节点，递归
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsRecursion(newHead.Next)
	newHead.Next = head
	return newHead

}

func swapPairsIteration(head *ListNode) *ListNode {

}

// swapPairs 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	firstSwap := true
	cur := head
	for cur != nil {
		if cur.Next != nil {
			// 保存原有的节点
			tmp := cur // 1
			tmpNext := &ListNode{Val: cur.Next.Val, Next: cur.Next.Next}
			fmt.Println(tmpNext, tmpNext.Next)

			// 交换当前节点和下一个节点
			cur := cur.Next // 1 变为 2
			// 修改指针指向
			cur.Next = tmp // 2 -> 1
			fmt.Println(tmpNext, tmpNext.Next)
			cur.Next.Next = tmpNext.Next
			fmt.Println(cur.Next.Next)
			if firstSwap {
				head = cur
				firstSwap = false
			}
			cur = cur.Next.Next
		} else {
			return head
		}
	}
	return head
}

// printLinkedList 打印链表
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
