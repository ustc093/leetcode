package binary_tree

import "container/list"

func InorderSuccessor(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	for ; stack.Len() != 0 || root != nil; {
		if root != nil {
			stack.PushBack(root)
			root = root.Left
		} else {
			ele := stack.Back()
			stack.Remove(ele)
			treeNode := ele.Value.(*TreeNode)

			if treeNode == node {
				// 找到后继节点

				if treeNode.Right == nil {
					// 没有右孩子 则返回父节点
					if stack.Len() == 0 {
						return nil
					} else {
						return stack.Back().Value.(*TreeNode)
					}
				} else {
					// 有右孩子的话，找到最左的节点
					return helperInorderSuccessor(treeNode.Right)
				}
			}
			root = treeNode.Right
		}
	}

	return nil
}

func helperInorderSuccessor(root *TreeNode) *TreeNode {
	curr := root
	for ; curr.Left != nil; {
		curr = curr.Left
	}

	return curr
}
