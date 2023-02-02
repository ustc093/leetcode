package hot_100

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。


参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true

链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
*/

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n <= 1 {
		return true
	}

	index := 0
	root := postorder[n-1]
	for i, num := range postorder {
		if num >= root {
			index = i
			break
		}
	}

	left := postorder[:index]
	right := postorder[index : n-1]

	for _, num := range left {
		if num > root {
			return false
		}
	}

	for _, num := range right {
		if num < root {
			return false
		}
	}

	return verifyPostorder(left) && verifyPostorder(right)
}
