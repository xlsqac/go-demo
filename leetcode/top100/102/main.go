// package main
// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&envId=top-100-liked
// 二叉树的层序遍历
// 广度优先遍历，一层一层遍历
package main

import (
	"container/list"
	"fmt"
)

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// main
func main() {
	{
		node1 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 9}
		node3 := &TreeNode{Val: 20}
		node4 := &TreeNode{Val: 15}
		node5 := &TreeNode{Val: 7}

		node1.Left = node2
		node1.Right = node3
		node3.Left = node4
		node3.Right = node5
		fmt.Println(levelOrder(node1))
	}

	{
		node1 := &TreeNode{Val: 1}
		fmt.Println(levelOrder(node1))
	}

	{
		fmt.Println(levelOrder(nil))

	}

}

// levelOrder 层序遍历
// 时间复杂度 O(N)
// 利用队列实现
func levelOrder(root *TreeNode) [][]int {
	queue := list.New()
	result := [][]int{}

	if root == nil {
		return result
	}
	queue.PushBack(root)

	for queue.Len() > 0 {
		n := queue.Len()
		vals := make([]int, n)
		for i := range vals {
			node := queue.Remove(queue.Front()).(*TreeNode)
			vals[i] = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, vals)
	}
	return result

}
