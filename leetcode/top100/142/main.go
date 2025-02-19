// package main
// https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked
// 环形链表 II
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	{
		node1 := &ListNode{Val: 3}
		node2 := &ListNode{Val: 2}
		node3 := &ListNode{Val: 0}
		node4 := &ListNode{Val: -4}

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node2
		fmt.Println(detectCycle(node1))
		fmt.Println(detectCycleSpace(node1))
	}
	{
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}

		node1.Next = node2
		node2.Next = node1
		fmt.Println(detectCycle(node1))
		fmt.Println(detectCycleSpace(node1))
	}
	{
		node1 := &ListNode{Val: 1}
		fmt.Println(detectCycle(node1))
		fmt.Println(detectCycleSpace(node1))
	}
}

// detectCycle 检测链表是否有环
// 时间复杂度 O(n)，空间复杂度 O(n)
// 通过遍历把每个节点存到 map 中，如果遍历到的节点在 map 中存在，则说明有环
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		} else {
			m[head] = struct{}{}
		}
		head = head.Next
	}

	return nil
}

// detectCycleSpace 检测链表是否有环
// 时间复杂度 O(n)，空间复杂度 O(1)
// 快慢指针，如果有环，快指针会追上慢指针
// 没有用哈希表来存储，节省了空间
func detectCycleSpace(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		// 如果有环，fast 会追上 slow
		if fast == slow {
			p := head
			// 找到环
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
