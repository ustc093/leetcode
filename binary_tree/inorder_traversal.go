package binary_tree

import "container/list"

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}
	InOrderTraversal(root.Left)
	res = append(res, root.Val)
	InOrderTraversal(root.Right)
	return res
}

func InOrderTraversalII(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	curr := root
	stack := list.New()

	for ; stack.Len() != 0 || curr != nil; {
		if curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		} else {
			ele := stack.Back()
			stack.Remove(ele)
			if node, ok := ele.Value.(*TreeNode); ok {
				res = append(res, node.Val)
				curr = node.Right
			}
		}
	}
	return res
}
