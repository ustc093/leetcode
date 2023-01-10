package hot_100

import "container/list"

/**
请实现一个函数按照之字形顺序打印二叉树，
即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
*/

func levelOrderIII(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	reverse := false
	for queue.Len() != 0 {
		n := queue.Len()
		row := make([]int, n)
		for i := 0; i < n; i++ {
			ele := queue.Front()
			tmp := ele.Value.(*TreeNode)
			if tmp != nil {
				if reverse {
					row[n-i-1] = tmp.Val
				} else {
					row[i] = tmp.Val
				}

				if tmp.Left != nil {
					queue.PushBack(tmp.Left)
				}

				if tmp.Right != nil {
					queue.PushBack(tmp.Right)
				}
			}
			queue.Remove(ele)

		}
		reverse = !reverse
		res = append(res, row)
	}

	return res
}
