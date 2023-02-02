package hot_100

/**
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

    3
   / \
  9  20
    /  \
   15   7

示例 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

示例 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]

链接：https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	rootVal := preorder[0]
	res := &TreeNode{
		Val: rootVal,
	}
	index := 0
	for i, num := range inorder {
		if num == rootVal {
			index = i
			break
		}
	}
	preLeft := preorder[1:index+1]
	preRight := preorder[index+1:]

	inLeft := inorder[:index]
	inright := inorder[index+1:]

	res.Left = buildTree(preLeft, inLeft)
	res.Right = buildTree(preRight, inright)

	return res
}
