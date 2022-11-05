package binary_tree

import "container/list"

var res []int

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
	return res
}

func PreOrderTraversalII(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for ; stack.Len() != 0; {
		element := stack.Back()
		if node, ok := element.Value.(*TreeNode); ok {
			res = append(res, node.Val)
			left := node.Left
			right := node.Right
			stack.Remove(element)
			if right != nil {
				stack.PushBack(right)
			}
			if left != nil {
				stack.PushBack(left)
			}
		}
	}

	return res
}
