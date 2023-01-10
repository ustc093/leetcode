package hot_100

import "container/list"

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
*/

func levelOrderII(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for ; queue.Len() != 0; {
		n := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			ele := queue.Front()
			node := ele.Value.(*TreeNode)
			if node != nil {
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}

				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
			queue.Remove(ele)
		}

		res = append(res, tmp)
	}

	return res
}
