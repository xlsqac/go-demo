// package main
// https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
// 回文链表，判断一个链表是否为回文链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	{
		n1 := &ListNode{Val: 1, Next: nil}
		n2 := &ListNode{Val: 2, Next: nil}
		n3 := &ListNode{Val: 2, Next: nil}
		n4 := &ListNode{Val: 1, Next: nil}

		n1.Next = n2
		n2.Next = n3
		n3.Next = n4

		fmt.Println(isPalindrome(n1), true)
		fmt.Println(isPalindromeSlice(n1), true)
	}

	{
		n1 := &ListNode{Val: 1, Next: nil}
		n2 := &ListNode{Val: 2, Next: nil}
		n1.Next = n2
		fmt.Println(isPalindrome(n1), false)
		fmt.Println(isPalindromeSlice(n1), false)
	}

	{
		n1 := &ListNode{Val: 1, Next: nil}
		n2 := &ListNode{Val: 1, Next: nil}
		n3 := &ListNode{Val: 2, Next: nil}
		n4 := &ListNode{Val: 1, Next: nil}

		n1.Next = n2
		n2.Next = n3
		n3.Next = n4

		fmt.Println(isPalindrome(n1), false)
		fmt.Println(isPalindromeSlice(n1), false)
	}
}

// 把链表转换成切片然后用双指针去做
func isPalindromeSlice(head *ListNode) bool {
	slice := liked2slice(head)
	pa, pb := 0, len(slice)-1
	for pa < pb {
		if slice[pa] != slice[pb] {
			return false
		}
		pa++
		pb--
	}
	return true

}

func liked2slice(head *ListNode) []int {
	slice := make([]int, 0)
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}

// isPalindrome 判断一个链表是否为回文链表
// 暴力思路，先反转再比较
// 时间复杂度 O(2n), 空间复杂度 O(n)
func isPalindrome(head *ListNode) bool {
	reversedHead := reverseList(head)
	for head != nil {
		fmt.Println(head.Val, reversedHead.Val)
		if head.Val != reversedHead.Val {
			return false
		}
		head = head.Next
		reversedHead = reversedHead.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		newNode := &ListNode{Val: head.Val}
		newNode.Next = newHead
		newHead = newNode
		head = head.Next
	}
	return newHead
}
