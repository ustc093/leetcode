package binary_tree

import "container/list"

const (
	UIntMIN = uint(0)
	UIntMAX = ^uint(0)
	IntMAX  = int(UIntMAX >> 1)
	IntMIN  = ^IntMAX
)

func JudgeSearchBinaryTree(head *TreeNode) bool {
	if head == nil {
		return true
	}
	stack := list.New()
	pre := IntMIN
	for ; stack.Len() != 0 || head != nil; {
		if head != nil {
			stack.PushBack(head)
			head = head.Left
		} else {
			ele := stack.Back()
			stack.Remove(ele)

			if node, ok := ele.Value.(*TreeNode); ok {
				if node.Val < pre {
					return false
				}
				pre = node.Val
				head = node.Right
			}

		}
	}

	return true
}

func JudgeSearchBinaryTreeII(head *TreeNode) bool {
	tmp := IntMIN
	return helper(head, &tmp)
}

func helper(head *TreeNode, preVal *int) bool {
	if head == nil {
		return true
	}
	l := helper(head.Left, preVal)

	if !l || head.Val < *preVal {
		return false
	}
	*preVal = head.Val
	return helper(head.Right, preVal)
}
