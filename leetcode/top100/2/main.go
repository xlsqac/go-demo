// package main
// https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked
// 两数相加
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	{
		node1 := &ListNode{Val: 2}
		node2 := &ListNode{Val: 4}
		node3 := &ListNode{Val: 3}
		node1.Next = node2
		node2.Next = node3

		node11 := &ListNode{Val: 5}
		node12 := &ListNode{Val: 6}
		node13 := &ListNode{Val: 4}
		node11.Next = node12
		node12.Next = node13

		node := addTwoNumbers(node1, node11)
		printListNode(node)
		fmt.Println()
	}
	{
		node1 := &ListNode{Val: 0}

		node11 := &ListNode{Val: 0}

		node := addTwoNumbers(node1, node11)
		printListNode(node)
		fmt.Println()
	}
	{
		node1 := &ListNode{Val: 9}
		node2 := &ListNode{Val: 9}
		node3 := &ListNode{Val: 9}
		node4 := &ListNode{Val: 9}
		node5 := &ListNode{Val: 9}
		node6 := &ListNode{Val: 9}
		node7 := &ListNode{Val: 9}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node5.Next = node6
		node6.Next = node7

		node11 := &ListNode{Val: 9}
		node12 := &ListNode{Val: 9}
		node13 := &ListNode{Val: 9}
		node14 := &ListNode{Val: 9}
		node11.Next = node12
		node12.Next = node13
		node13.Next = node14
		node := addTwoNumbers(node1, node11)
		printListNode(node)
	}
}

// addTwoNumbers 两数相加
// 时间复杂度 O(max(m, n))，空间复杂度 O(1)
// 依次遍历两个链表，相加，如果有进位则加1，如果和大于等于10，则减去10，将结果放入新链表中
// 如果两个链表长度不一样，短的链表后面的值为0，还要注意最后一位相加后有进位的情况
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addFlag := 0
	nodeCount := 0
	node := &ListNode{}
	head := node
	for l1 != nil || l2 != nil || addFlag == 1 {
		l1Val := 0
		if l1 == nil {
			l1Val = 0
			l1 = nil
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}
		l2Val := 0
		if l2 == nil {
			l2Val = 0
			l2 = nil
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}
		num := l1Val + l2Val + addFlag
		if num >= 10 {
			num = num - 10
			addFlag = 1
		} else {
			addFlag = 0
		}

		if nodeCount == 0 {
			node.Val = num
		} else {
			node.Next = &ListNode{Val: num}
			node = node.Next
		}

		nodeCount++
	}
	return head
}

func printListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}
