package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	originArr              = []int{1, 2, 3, 4, 5, 6}
	preOrderTraversalArr   = []int{1, 2, 4, 5, 3, 6}
	inOrderTraversalArr   = []int{4, 2, 5, 1, 6, 3}
	postOrderTraversalArr = []int{4, 5, 2, 6, 3, 1}
)

func TestPreOrderTraversal(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := PreOrderTraversal(root)
	assert.Equal(t, preOrderTraversalArr, res)

	res = PreOrderTraversalII(root)
	assert.Equal(t, preOrderTraversalArr, res)
}

func TestInOrderTraversal(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := InOrderTraversal(root)
	assert.Equal(t, inOrderTraversalArr, res)

	res = InOrderTraversalII(root)
	assert.Equal(t, inOrderTraversalArr, res)
}

func TestPostOrderTraversal(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := PostOrderTraversal(root)
	assert.Equal(t, postOrderTraversalArr, res)

	res = PostOrderTraversalII(root)
	assert.Equal(t, postOrderTraversalArr, res)

}

func TestLevelOrder(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := LevelOrder(root)
	print(res)
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
