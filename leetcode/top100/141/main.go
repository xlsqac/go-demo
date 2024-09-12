// package main
// https://leetcode.cn/problems/linked-list-cycle/?envType=study-plan-v2&envId=top-100-liked
// 环形链表，检测链表是否有环
package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// main
func main() {
	{
		n1 := &ListNode{Val: 3, Next: nil}
		n2 := &ListNode{Val: 2, Next: nil}
		n3 := &ListNode{Val: 0, Next: nil}
		n4 := &ListNode{Val: -4, Next: nil}

		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n2
		fmt.Println(hasCycle(n1), true)
	}

	{
		n1 := &ListNode{Val: 1, Next: nil}
		n2 := &ListNode{Val: 2, Next: nil}

		n1.Next = n2
		n2.Next = n1
		fmt.Println(hasCycle(n1), true)
	}

	{
		n1 := &ListNode{Val: 1, Next: nil}
		fmt.Println(hasCycle(n1), true)
	}
}

// hasCycel 判断链表是否有环
// 利用哈希表解决，如果有环，那么一定会出现重复的节点
// 时间复杂度 O(n)，空间复杂度 O(n)
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}
