// 二叉搜索树
// 查找、插入和删除 O(logn)
// 中序遍历

package tree

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

type binarySearchTree struct {
	root *treeNode
}

func newTreeNode(num int) *treeNode {
	return &treeNode{Val: num}
}

func newBinarySearchTree(num int) *binarySearchTree {
	node := newTreeNode(num)
	return &binarySearchTree{
		root: node,
	}
}

// 查找
// O(logn)
func (bst *binarySearchTree) search(num int) *treeNode {
	node := bst.root
	for node != nil {
		if node.Val < num {
			node = node.Right
		} else if node.Val > num {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

// 插入
// O(logn)
func (bst *binarySearchTree) insert(num int) {
	if bst.root == nil {
		bst.root = newTreeNode(num)
		return
	}
	cur := bst.root

	// 借助 pre 来保存插入节点的父节点
	var pre *treeNode = nil
	for cur != nil {
		// bst 不允许有重复的结点，如果有的话不执行插入
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else if cur.Val > num {
			cur = cur.Left
		}
	}
	// 插入节点
	node := newTreeNode(num)
	if pre.Val < num {
		pre.Right = node
	} else {
		pre.Left = node
	}

}

// 删除
// O(logn)
func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	// 待删除节点的父节点
	var pre *treeNode = nil
	// 先找到需要删除的节点
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else if cur.Val > num {
			cur = cur.Left
		}
	}
	// 根据待删除节点的度的不同需要处理不同的逻辑
	// cur 就是要删除的节点
	// 如果度是 0或1时，把 cur 直接替换成 cur.left 或 cur.right
	// 如果度是 2，需要用右子树的最小节点或左子树的最大节点来替换
	// 替换时需要把替换的节点删掉，然后替换的值赋给要删除的节点
	if cur.Right == nil || cur.Left == nil {
		var child *treeNode = nil
		// 把子节点取出来赋值给父节点
		// 把有值的那个子节点取出来
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		// 把子节点送给父节点
		if pre.Left == cur {
			pre.Left = child
		} else {
			pre.Right = child
		}
	} else {
		// 找到中序遍历中待删除节点的下一个节点
		// 这里选右子树的最小节点, 也就是中序遍历的下一个结点
		tmp := cur.Right
		for cur.Left != nil {
			tmp = cur.Left
		}
		// 删除这个节点，之后把该节点的值赋给待删除结点
		// 这里只能改值，不能改整个对象，因为对象中保存有指针
		bst.remove(tmp.Val)
		cur.Val = tmp.Val
	}
}
