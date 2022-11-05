package binary_tree

import (
	"container/list"
)

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	res = append(res, root.Val)
	return res
}

func PostOrderTraversalII(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	tmpStack := list.New()
	stack := list.New()
	curr := root
	tmpStack.PushBack(curr)
	for ; tmpStack.Len() != 0; {
		ele := tmpStack.Back()
		tmpStack.Remove(ele)

		if node,ok := ele.Value.(*TreeNode);ok {
			stack.PushBack(node)
			left := node.Left
			right := node.Right
			if left != nil {
				tmpStack.PushBack(left)
			}
			if right != nil {
				tmpStack.PushBack(right)
			}
		}
	}

	for ; stack.Len() != 0; {
		ele := stack.Back()
		stack.Remove(ele)

		if node,ok := ele.Value.(*TreeNode);ok {
			res = append(res, node.Val)
		}

	}
	return res
}
