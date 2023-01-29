package hot_100

import "testing"

func TestPathSumHelp(t *testing.T) {
	sum := pathSum(genBinaryTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 5, 1}, 0), 22)
	println(sum)
}

func genBinaryTree(arr []int, index int) *TreeNode {
	if index >= len(arr) || arr[index] == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[index]}
	root.Left = genBinaryTree(arr, index*2+1)
	root.Right = genBinaryTree(arr, index*2+2)
	return root
}
