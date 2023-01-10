package hot_100

import "container/list"

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树:[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回：
[3,9,20,15,7]

链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
*/

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := make([]int,0)

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for ;queue.Len() != 0; {
		ele := queue.Front()
		node := ele.Value.(*TreeNode)
		if node != nil {
			res = append(res, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		queue.Remove(ele)
	}
	return res
}
