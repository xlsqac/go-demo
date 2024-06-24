// Package tree
// 基于数组表示的二叉树
// 基于数组表示时，所有的树都认为是完美二叉树
// 实际上没有的节点用 nil 来表示
// 对于索引为 i 的结点来说，他的左节点的索引就是 2i + 1，右节点的索引就是 2i + 2
package tree

// 数组表示的二叉树
type arrayBinaryTree struct {
	tree []any
}

// 构造函数
func newArrayBinaryTree(arr []any) *arrayBinaryTree {
	return &arrayBinaryTree{tree: arr}
}

// 列表容量
func (a *arrayBinaryTree) size() int {
	return len(a.tree)
}

// 获取索引为 i 节点的值
func (a *arrayBinaryTree) val(i int) any {
	return a.tree[i]
}

// 获取索引为 i 节点的左子节点的索引
func (a *arrayBinaryTree) left(i int) int {
	return 2*i + 1
}

// 获取索引为 i 节点的右子节点的索引
func (a *arrayBinaryTree) right(i int) int {
	return 2*i + 2
}

// 获取索引为 i 节点的父节点的索引
func (a *arrayBinaryTree) partner(i int) int {
	// 索引为 0 的结点就是根节点
	//if i == 0 {
	//	return nil
	//}
	//if i%2 == 0 {
	//	return (i - 2) / 2
	//} else {
	//	return (i - 1) / 2
	//}
	// 利用整数相除返回的也是整数的特性直接简写成
	return (i - 1) / 2
}

// 层序遍历
func (a *arrayBinaryTree) levelOrder() []any {
	// 插入的数组本身就是层序遍历好的
	// 把不存在的节点去掉，也就是 nil
	var res []any
	for i := 0; i < a.size(); i++ {
		if a.val(i) != nil {
			res = append(res, a.val(i))
		}
	}
	return res
}

// 深度优先遍历
func (a *arrayBinaryTree) dfs(i int, order string, res *[]any) {
	if a.val(i) == nil {
		return
	}
	// 前序遍历
	if order == "pre" {
		*res = append(*res, a.val(i))
	}
	a.dfs(a.left(i), order, res)
	// 中序遍历
	if order == "in" {
		*res = append(*res, a.val(i))
	}
	a.dfs(a.right(i), order, res)
	// 后序遍历
	if order == "post" {
		*res = append(*res, a.val(i))
	}
}

// 前序遍历
func (a *arrayBinaryTree) preOrder() []any {
	var res []any
	a.dfs(0, "pre", &res)
	return res
}

// 中序遍历
func (a *arrayBinaryTree) inOrder() []any {
	var res []any
	a.dfs(0, "in", &res)
	return res
}

// 后序遍历
func (a *arrayBinaryTree) postOrder() []any {
	var res []any
	a.dfs(0, "post", &res)
	return res
}
