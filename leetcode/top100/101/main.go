// package main
// https://leetcode.cn/problems/symmetric-tree/?envType=study-plan-v2&envId=top-100-liked
// 判断一个树是不是对称二叉树
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
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 3}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	fmt.Println(isSymmetricWithIteration(node1))

	{
		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 2}
		node3 := &TreeNode{Val: 2}
		node4 := &TreeNode{Val: 3}
		node5 := &TreeNode{Val: 3}

		node1.Left = node2
		node1.Right = node3

		node2.Right = node4
		node3.Right = node5
		fmt.Println(isSymmetricWithIteration(node1))
	}
}

// check 判断是否对称二叉树
// 时间复杂度：O(n)
func check(p, q *TreeNode) bool {
	// 如果 p 和 q 都为 nil，则表示当前节点对称，返回 true
	if p == nil && q == nil {
		return true
	}
	// 都为 nil 在上面判断了这里如果能走进去说明只有一个为 nil，说明不是对称二叉树，返回 false
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// isSymmetric 判断是否是对称二叉树
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

// isSymmetricWithIteration 使用迭代的方法判断是否是对称二叉树
// 时间复杂度：O(n)
func isSymmetricWithIteration(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		node1 := queue.Remove(queue.Front()).(*TreeNode)
		node2 := queue.Remove(queue.Front()).(*TreeNode)
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue.PushBack(node1.Left)
		queue.PushBack(node2.Right)

		queue.PushBack(node1.Right)
		queue.PushBack(node2.Left)

	}
	return true
}
