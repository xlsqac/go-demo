// package main
// https://leetcode.cn/problems/binary-tree-inorder-traversal/?envType=study-plan-v2&envId=top-100-liked
// 二叉树的中序遍历
// 中序遍历是深度优先遍历（dfs）的一种，体现为从左往右遍历
package main

import "fmt"

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// main
func main() {
	n1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 3, Left: nil, Right: nil}

	n1.Right = n2
	n2.Left = n3
	s := inorderTraversal(n1)
	fmt.Println(s, []int{1, 3, 2})
}

// inorderTraversal
func inorderTraversal(root *TreeNode) []int {
	var nums []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return nums
}
