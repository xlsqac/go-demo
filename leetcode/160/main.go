// package main
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=problem-list-v2&envId=2cktkvj
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{4, 1, 8, 4, 5}
	b := []int{5, 6, 1, 8, 4, 5}

	headA := slice2LinkedList(a)
	headB := slice2LinkedList(b)

	node := getIntersectionNode(headA, headB)
	if node == nil {
		fmt.Println(nil)
	} else {
		fmt.Println(node.Val)
	}
}

func slice2LinkedList(s []int) *ListNode {
	head := &ListNode{Val: s[0], Next: nil}
	prev := head
	for _, v := range s[1:] {
		node := &ListNode{Val: v, Next: nil}
		prev.Next = node
		prev = node
	}
	return head
}

// getIntersectionNodeOnm 暴力解法 O(n*m)
func getIntersectionNodeOnm(headA, headB *ListNode) *ListNode {
	for headA != nil {
		hb := headB
		for hb != nil {
			if headA == hb {
				return hb
			}
			hb = hb.Next
		}
		headA = headA.Next
	}
	return nil
}

// getIntersectionNode 官网解法，利用哈希
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapA := make(map[*ListNode]bool)
	for headA != nil {
		mapA[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if mapA[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
