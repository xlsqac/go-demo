// package reverse
package reverse

import "xlsqac/interview/algo/linked/listnode"

func Reverse(head *listnode.ListNode) *listnode.ListNode {
	var prev *listnode.ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next // 先暂存下一个节点
		curr.Next = prev  // 当前节点的下一个节点指向 prev
		prev = curr       // prev 移动到当前节点
		curr = next       // 当前节点移动到下一个节点
	}
	return prev

}
