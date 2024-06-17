// package main
// https://leetcode.cn/problems/linked-list-cycle/description/
package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    slow := head
    fast := head.Next
    var prev *ListNode

    for slow != fast {
        if fast == nil || fast.Next == nil {
            return nil
        }
        fmt.Println(slow.Val, fast.Val)
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    return prev

}

func main() {
    //s := []int{3, 2, 0, -4}
    node4 := &ListNode{Val: -4}
    node3 := &ListNode{Val: 0, Next: node4}
    node2 := &ListNode{Val: 2, Next: node3}
    node1 := &ListNode{Val: 3, Next: node2}
    node4.Next = node2 // 形成环
    fmt.Println(detectCycle(node1))

}
