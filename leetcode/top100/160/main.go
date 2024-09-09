// package main
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
// 链表相交，找到并返回两个相交节点的起始节点
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := &ListNode{Val: 4, Next: nil}
	a2 := &ListNode{Val: 1, Next: nil}

	c1 := &ListNode{Val: 8, Next: nil}
	c2 := &ListNode{Val: 4, Next: nil}
	c3 := &ListNode{Val: 5, Next: nil}

	b1 := &ListNode{Val: 5, Next: nil}
	b2 := &ListNode{Val: 6, Next: nil}
	b3 := &ListNode{Val: 1, Next: nil}

	c1.Next = c2
	c2.Next = c3

	a1.Next = a2
	a2.Next = c1

	b1.Next = b2
	b2.Next = b3
	b3.Next = c1

	intersectNode := getIntersectionNode(a1, b1)
	fmt.Println(intersectNode, c1.Val)

	{
		intersectNode := getIntersectionNodeHash(a1, b1)
		fmt.Println(intersectNode, c1.Val)
	}

	{
		intersectNode := getIntersectionNodeDoublePointer(a1, b1)
		fmt.Println(intersectNode, c1.Val)
	}

}

// getIntersectionNode
// 暴力法，O(n*m)，依次遍历 a 和 b
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		currB := headB
		for currB != nil {
			if currB == headA {
				return currB
			}
			currB = currB.Next
		}
		headA = headA.Next
	}
	return nil
}

// getIntersectionNodeHash
// 利用哈希，先遍历链表 A，把结果放到哈希中，再遍历链表 B，判断是否在哈希中
// 时间复杂度 O(n+m)，空间复杂度 O(n)
func getIntersectionNodeHash(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	h := make(map[*ListNode]bool)
	for headA != nil {
		h[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if h[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// getIntersectionNodeDoublePointer
// 双指针法，设定 2 个指针，一个从 A 出发，一个从 B 出发
// A 走完了，再从 B 开始走，B 走完了，从 A 开始走
// 如果有交点，那么一定会相遇，因为走过的路程是一样的
// 最后返回 pA 或者 pB，如果有值那么就说明相交了，没有值则不相交
// 判断条件是 pA!= pB，因为如果有交点，那么一定会相遇，相遇的点就是交点，没值则都为 nil
func getIntersectionNodeDoublePointer(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
