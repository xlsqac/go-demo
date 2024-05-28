package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	slow.Next = fast
	return head

}

func main() {
	a := ListNode{
		Val: 1,
	}

	b := ListNode{
		Val: 1,
	}

	c := ListNode{
		Val: 2,
	}

	a.Next = &b
	b.Next = &c

	deleteDuplicates(&a)
}
