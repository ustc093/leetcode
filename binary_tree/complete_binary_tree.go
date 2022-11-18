package binary_tree

import list "container/list"

// JudgeCompleteBinaryTree 层序遍历
func JudgeCompleteBinaryTree(head *TreeNode) bool {
	if head == nil {
		return true
	}

	// 叶子节点标志
	flag := false

	queue := list.New()
	queue.PushBack(head)

	for ; queue.Len() != 0; {
		n := queue.Len()
		for i := 0; i < n; i++ {
			ele := queue.Front()
			queue.Remove(ele)
			if node, ok := ele.Value.(*TreeNode); ok {
				if node.Left == nil && node.Right != nil {
					return false
				}
				if flag && (node.Left != nil || node.Right != nil) {
					return false
				}

				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}

				// 子孩子有一个不为空，以后节点的孩子都必须为空
				if node.Left == nil || node.Right == nil {
					flag = true
				}
			}
		}
	}
	return true
}
