// package main
// https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 二叉树的直径
package main

import (
	"fmt"
)

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// main
func main() {
	{
		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 2}
		node3 := &TreeNode{Val: 3}
		node4 := &TreeNode{Val: 4}
		node5 := &TreeNode{Val: 5}

		node1.Left = node2
		node1.Right = node3
		node2.Left = node4
		node2.Right = node5

		fmt.Println(diameterOfBinaryTree(node1))
	}

	{
		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 2}

		node1.Left = node2

		fmt.Println(diameterOfBinaryTree(node1))
	}
}

// diameterOfBinaryTree 返回二叉树的直径
// 时间复杂度 O(N)
// 直径：树中任意两个节点的路径长度
// 左子树的深度 + 右子树的深度 +1
func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0

	dfs = func(node *TreeNode) int {
		// 空节点，深度为0
		if node == nil {
			return 0
		}
		// 左儿子为根的子树的深度
		lLen := dfs(node.Left) // 1 2 4 nil return 0
		// 右儿子为根的子树的深度
		rLen := dfs(node.Right) // nil(4) return 0
		// 计算深度的最大值
		ans = max(ans, lLen+rLen) // 0
		// 返回节点的深度，左右子树深度的最大值 + 1
		// 这个结果经过 return 后，会作为左儿子或右儿子为根的子树的深度
		return max(lLen, rLen) + 1 // 1(4)
	}
	dfs(root)
	return ans
}
