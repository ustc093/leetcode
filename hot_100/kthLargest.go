package hot_100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count []int
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	help(root)
	return count[k-1]
}

func help(root *TreeNode) {
	if root == nil {
		return
	}
	help(root.Left)
	count = append(count,root.Val)
	help(root.Right)
}
