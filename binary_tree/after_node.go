package binary_tree

import "container/list"

// InorderSuccessor
// 题目给定一个二叉搜索树的根节点root和树中的某一个节点node，要求出节点node的中序遍历的后继Successor
// 所谓后继就是树所有值大于p节点值的节点中值最小的那个节点。如果后继不存在就返回空节点。
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
