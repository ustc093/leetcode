package binary_tree

/*
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetcod/
 * 最低公共祖先
 */

func LowestCommonAncestor(root, node1, node2 *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node1 || root == node2 {
		return root
	}

	l := LowestCommonAncestor(root.Left, node1, node2)
	r := LowestCommonAncestor(root.Right, node1, node2)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r
}
