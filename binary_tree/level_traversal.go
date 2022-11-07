package binary_tree

import "container/list"

func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for ; queue.Len() != 0; {
		tmp := make([]int, 0)

		n := queue.Len()
		for i := 0; i < n; i++ {
			// pop
			curr := queue.Front()
			queue.Remove(curr)

			if node, ok := curr.Value.(*TreeNode); ok {
				tmp = append(tmp, node.Val)
				left := node.Left
				right := node.Right
				if left != nil {
					queue.PushBack(left)
				}
				if right != nil {
					queue.PushBack(right)
				}
			}
		}

		res = append(res, [][]int{
			tmp,
		}...)
	}

	return res
}
