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
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。


示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
		 5
	   /   \
	  4      8
	 /      /  \
	11     13   4
   /  \        /  \
  7    2      5    1
输出：[[5,4,11,2],[5,8,4,5]]

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：[]

示例 3：
输入：root = [1,2], targetSum = 0
输出：[]

链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
*/

var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	pathSumHelp(root, target, make([]int, 0))
	return res
}

func pathSumHelp(root *TreeNode, target int, tmp []int) {
	if root == nil || root.Val > target {
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == target {
		newTmp := make([]int, len(tmp))
		copy(newTmp, tmp)
		res = append(res, newTmp)
	}

	pathSumHelp(root.Left, target-root.Val, tmp)
	pathSumHelp(root.Right, target-root.Val, tmp)
}
