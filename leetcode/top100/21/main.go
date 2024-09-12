// package main
// https://leetcode.cn/problems/merge-two-sorted-lists/?envType=study-plan-v2&envId=top-100-liked
// 合并两个有序链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// main
func main() {
	{
		n1 := &ListNode{Val: 1, Next: nil}
		n2 := &ListNode{Val: 2, Next: nil}
		n3 := &ListNode{Val: 4, Next: nil}

		n1.Next = n2
		n2.Next = n3

		n4 := &ListNode{Val: 1, Next: nil}
		n5 := &ListNode{Val: 3, Next: nil}
		n6 := &ListNode{Val: 4, Next: nil}
		n4.Next = n5
		n5.Next = n6

		l := mergeTwoLists(n1, n4)
		printList(l)
	}

	{
		var n1 *ListNode = nil
		var n2 *ListNode = nil
		l := mergeTwoLists(n1, n2)
		printList(l)
	}
	{
		var n1 *ListNode = nil
		n2 := &ListNode{Val: 0, Next: nil}
		l := mergeTwoLists(n1, n2)
		printList(l)
	}
}

// mergeTwoLists
// 遍历两个链表，将较小的节点添加到新链表中
// 时间复杂度是 O(m+n)，空间复杂度是 O(1)
// 遍历的方法代码有点烂
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 如果有一个链表为空，直接返回另一个链表
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	prevHead := &ListNode{Val: -1, Next: nil}
	prev := prevHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	// list1 和 list2 最多有一个不为空，将不为空的链表添加到新链表中
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return prevHead.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
