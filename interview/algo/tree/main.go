// package main
package main

import (
	"xlsqac/interview/algo/tree/levelordertraversal"
	"xlsqac/interview/algo/tree/treenode"
)

func main() {
	n1 := &treenode.TreeNode{Val: 1, Left: nil, Right: nil}
	n2 := &treenode.TreeNode{Val: 2, Left: nil, Right: nil}
	n3 := &treenode.TreeNode{Val: 3, Left: nil, Right: nil}
	n4 := &treenode.TreeNode{Val: 4, Left: nil, Right: nil}
	n5 := &treenode.TreeNode{Val: 5, Left: nil, Right: nil}
	n6 := &treenode.TreeNode{Val: 6, Left: nil, Right: nil}
	n7 := &treenode.TreeNode{Val: 7, Left: nil, Right: nil}

	// 组装树
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7

	levelordertraversal.LevalOrderTraversal(n1)
	levelordertraversal.LevalOrderTraversal1(n1)

}
