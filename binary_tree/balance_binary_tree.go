package binary_tree

import "math"

func JudgeBalanceBinaryTree(head *TreeNode) bool {
	res,_ := helperBalanceBinaryTree(head)
	return res
}

//
func helperBalanceBinaryTree(head *TreeNode) (bool, float64) {
	if head == nil {
		return true, 0
	}

	r, rHigh := helperBalanceBinaryTree(head.Left)
	l, lHigh := helperBalanceBinaryTree(head.Right)
	high := math.Max(rHigh, lHigh) + 1
	if math.Abs(rHigh-lHigh) > 1 {
		return false, high
	}
	return r && l, high
}
