// package levelordertraversal
// 二叉树的层序遍历
package levelordertraversal

import (
	"container/list"
	"fmt"
	"xlsqac/interview/algo/tree/treenode"
)

// LevalOrderTraversal 二叉树的层序遍历
// 遍历思路是一层一层遍历
// 先遍历根节点，再遍历根节点的左右子节点，在遍历根节点的左右子节点的左右子节点
// 利用队列的先进先出的思想来做 queue 保存节点的队列，nums 保存值的切片
// 先把根节点加到队列中，然后开启循环
func LevalOrderTraversal(root *treenode.TreeNode) {
	queue := list.New()
	queue.PushBack(root)

	nums := make([]int, 0)

	for queue.Len() > 0 {
		// remove 出来的是 any
		// 从头部弹出一个 node，第一次是根节点
		node := queue.Remove(queue.Front()).(*treenode.TreeNode)
		// 把弹出节点的值保存起来
		nums = append(nums, node.Val)

		// 判断弹出节点的左右子节点是否为空，不为空加到队尾
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}

	}
	fmt.Println(nums)
}

// LevalOrderTraversal1 二叉树的层序遍历修改版，如果上一层从左到右遍历，那么下一层就从右到左，反之亦然
// 给定范围是 2 层(3层)
// 1 2 3 7 6 5 4
func LevalOrderTraversal1(root *treenode.TreeNode) {
	queue := list.New()
	queue.PushBack(root)

	nums := make([]int, 0)

	counter := 0

	for queue.Len() > 0 {
		// remove 出来的是 any
		// 从头部弹出一个 node，第一次是根节点
		var node *treenode.TreeNode
		// 第一层是 2 个节点
		if counter < 3 {
			node = queue.Remove(queue.Front()).(*treenode.TreeNode)
		} else {
			node = queue.Remove(queue.Back()).(*treenode.TreeNode)
		}

		// 把弹出节点的值保存起来
		nums = append(nums, node.Val)

		// 判断弹出节点的左右子节点是否为空，不为空加到队尾
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		counter += 1

	}
	fmt.Println(nums)
}
