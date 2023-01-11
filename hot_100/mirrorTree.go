package hot_100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

  4
 /  \
 2   7
/ \  / \
1  3 6  9
镜像输出：

  4
 /  \
 7   2
/ \  / \
9  6 3 1


示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

链接：https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof
*/

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	right := root.Right
	left := root.Left
	root.Left = mirrorTree(right)
	root.Right = mirrorTree(left)

	return root
}
