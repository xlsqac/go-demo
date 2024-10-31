// package main
// https://leetcode.cn/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 翻转二叉树
package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 6}
	node6 := &TreeNode{Val: 9}

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	fmt.Println(levelOrder(root))

	invertedTreeRootNode := invertTree(root)

	// 输出翻转后的二叉树
	fmt.Println(levelOrder(invertedTreeRootNode))

	{
		root := &TreeNode{Val: 2}
		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 3}

		root.Left = node1
		root.Right = node2
		fmt.Println(levelOrder(root))

		invertedTreeRootNode := invertTree(root)

		// 输出翻转后的二叉树
		fmt.Println(levelOrder(invertedTreeRootNode))
	}
}

// invertTree 翻转二叉树 时间复杂度 O(n)，利用递归实现
// 先交换当前节点的左右子树，然后递归交换左右子树的左右子树
// 当前节点的左右节点都为 nil 时，停止递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	curr := root
	curr.Left, curr.Right = curr.Right, curr.Left
	if curr.Left == nil && curr.Right == nil {
		return root
	}
	invertTree(curr.Left)
	invertTree(curr.Right)
	return root
}

// levelOrder 层序遍历，把树遍历后加入到切片中
// 层序遍历是广度优先遍历，一般利用队列实现
func levelOrder(root *TreeNode) []int {
	queue := list.New()
	queue.PushBack(root)
	nums := []int{}
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}
