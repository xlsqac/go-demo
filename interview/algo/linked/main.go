// package main
package main

import (
	"fmt"
	"xlsqac/interview/algo/linked/listnode"
	"xlsqac/interview/algo/linked/reverse"
)

func main() {
	n0 := &listnode.ListNode{Val: 0, Next: nil}
	n1 := &listnode.ListNode{Val: 1, Next: nil}
	n2 := &listnode.ListNode{Val: 2, Next: nil}
	n3 := &listnode.ListNode{Val: 3, Next: nil}
	n4 := &listnode.ListNode{Val: 4, Next: nil}

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	printLinked(n0)

	fmt.Println("reverse linked")
	head := reverse.Reverse(n0)
	printLinked(head)
}

// printLinked 打印一个链表
func printLinked(head *listnode.ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
