// package main
// https://leetcode.cn/problems/linked-list-cycle/description/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true

}

func main() {
	//s := []int{3, 2, 0, -4}
	node4 := &ListNode{Val: -4}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 3, Next: node2}
	node4.Next = node2 // 形成环
	fmt.Println(hasCycle(node1))

}
