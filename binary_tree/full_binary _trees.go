package binary_tree

import "math"

func JudgeFullBinaryTree(head *TreeNode) bool {
	if head == nil {
		return false
	}
	high, count := helperFullBinaryTree(head)
	return math.Pow(2, high)-1 == count
}

func helperFullBinaryTree(head *TreeNode) (float64, float64) {
	if head == nil {
		return 0, 0
	}

	lHigh, lCount := helperFullBinaryTree(head.Left)
	rHigh, rCount := helperFullBinaryTree(head.Right)

	high := math.Max(lHigh, rHigh) + 1
	count := lCount + rCount + 1
	return high, count
}
