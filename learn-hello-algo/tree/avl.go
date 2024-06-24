// package tree
// avl 树，平衡搜索二叉搜索树
// 当 bst 不平衡时会通过左旋、右旋以及其组合来重新实现平衡
// 插入、删除、查找
// 左旋、右旋、先左旋后右旋和先右旋后左旋
// 左子树失衡需要右旋，右子树失衡需要左旋

package tree

// TreeNode 树结点
type TreeNode struct {
	Val    int       // 节点值
	Height int       // 节点高度
	Left   *TreeNode // 左子节点引用
	Right  *TreeNode // 右子节点引用
}

// aVLTree val 树
type aVLTree struct {
	root *TreeNode
}

// NewTreeNode 实例化一个结点
func NewTreeNode(num int) *TreeNode {
	return &TreeNode{Val: num}
}

// height 获取节点，叶节点高度为 0，空节点高度为 -1
func (t *aVLTree) height(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

// updateHeight 更新节点高度
// 高度等于 1 + 该节点的子节点最大高度
func (t *aVLTree) updateHeight(node *TreeNode) {
	lh := t.height(node.Left)
	lr := t.height(node.Right)

	node.Height = 1 + max(lh, lr)
}

// balanceFactor 获取节点平衡因子
func (t *aVLTree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// f = 左子树高度 - 右子树高度
	// -1 <= f <= 1
	// 一颗 avl 的任意节点的平衡因子都满足以上公式
	return t.height(node.Left) - t.height(node.Right)
}

// rightRotate 右旋
// 失衡节点为 node，左子节点为 child
// partner -> node -> child 变为 partner -> child -> node
// 如果 child 有右子节点，则该右子节点成为 node 的左子节点
func (t *aVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	child.Right = node
	node.Left = grandChild
	t.updateHeight(child)
	t.updateHeight(node)
	return child
}

// leftRotate 左旋
// 和右旋镜像即可
func (t *aVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	child.Left = node
	node.Right = grandChild
	t.updateHeight(child)
	t.updateHeight(node)
	return child
}

// rotate 旋转
// 根据失衡节点的平衡因子以及较高一侧子节点的平衡因子的正负号确定如何旋转
// 失衡节点的平衡因子 > 1，属于左偏树，右旋
// 失衡节点的平衡因子 < -1，属于右偏树，左旋
// 左偏树子节点的平衡因子 < 0 先左旋再右旋，>= 0 时只需右旋
// 右偏树子节点的平衡因子 > 0 先右旋再左旋，<= 0 时只需左旋
func (t *aVLTree) rotate(node *TreeNode) *TreeNode {
	bf := t.balanceFactor(node)
	if bf > 1 {
		// 左偏树
		// 左子节点小于 0 时需要先把左子节点左旋再右旋失衡节点
		if t.balanceFactor(node.Left) < 0 {
			node.Left = t.leftRotate(node.Left)
		}
		return t.rightRotate(node)
	} else if bf < -1 {
		// 右偏树
		// 右子节点小于 0 时需要先把右子节点右旋再左旋失衡节点
		if t.balanceFactor(node.Right) > 0 {
			node.Right = t.rightRotate(node.Right)
		}
		return t.leftRotate(node)

	}
	return node
}

// insert 插入节点
func (t *aVLTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

func (t *aVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}
	t.updateHeight(node)
	t.rotate(node)
	return node
}

// remove 删除节点
func (t *aVLTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

func (t *aVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return node
	}
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				return nil
			} else {
				node = child
			}
		} else {
			temp := node.Right
			if temp != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, node.Val)
			node.Val = temp.Val
		}
	}
	t.updateHeight(node)
	t.rotate(node)
	return node
}

// search 查找结点
func (t *aVLTree) search(val int) *TreeNode {
	node := t.root
	for node != nil {
		if node.Val > val {
			node = node.Left
		} else if node.Val < val {
			node = node.Right
		} else {
			break
		}
	}
	return node
}
