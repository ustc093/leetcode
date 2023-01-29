package hot_100

import "testing"

func TestKthLargest(t *testing.T) {
	kthLargest(genBinaryTree([]int{3, 1, 4, 0, 2}, 0), 1)
}
