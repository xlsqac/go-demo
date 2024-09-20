// package main
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 104. 二叉树的最大深度
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 20, Left: nil, Right: nil}
	n4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	n5 := &TreeNode{Val: 7, Left: nil, Right: nil}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	fmt.Println(maxDepth(n1), 3)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
