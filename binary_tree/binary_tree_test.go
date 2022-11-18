package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	originArr             = []int{1, 2, 3, 4, 5, 6}
	preOrderTraversalArr  = []int{1, 2, 4, 5, 3, 6}
	inOrderTraversalArr   = []int{4, 2, 5, 1, 6, 3}
	postOrderTraversalArr = []int{4, 5, 2, 6, 3, 1}

	searchBinaryTree     = []int{6, 5, 10, 4, 0, 7, 11}
	notBalanceBinaryTree = []int{6, 5, 3, 4, 0, 0, 0, 1}
	fullBinaryTree       = []int{1, 2, 3, 4, 5, 6, 7}
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

func TestJudgeSearchBinaryTree(t *testing.T) {
	head := genBinaryTree(searchBinaryTree, 0)
	res := JudgeSearchBinaryTree(head)
	assert.True(t, res)

	head = genBinaryTree(originArr, 0)
	res = JudgeSearchBinaryTree(head)
	assert.False(t, res)

	head = genBinaryTree(searchBinaryTree, 0)
	res = JudgeSearchBinaryTreeII(head)
	assert.True(t, res)

	head = genBinaryTree(originArr, 0)
	res = JudgeSearchBinaryTreeII(head)
	assert.False(t, res)
}

func TestJudgeCompleteBinaryTree(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := JudgeCompleteBinaryTree(root)
	assert.True(t, res)

	root = genBinaryTree(searchBinaryTree, 0)
	res = JudgeCompleteBinaryTree(root)
	assert.False(t, res)
}

func TestJudgeBalanceBinaryTree(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := JudgeCompleteBinaryTree(root)
	assert.True(t, res)

	root = genBinaryTree(notBalanceBinaryTree, 0)
	res = JudgeCompleteBinaryTree(root)
	assert.False(t, res)
}

func TestJudgeFullBinaryTree(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := JudgeFullBinaryTree(root)
	assert.False(t, res)

	root = genBinaryTree(fullBinaryTree, 0)
	res = JudgeFullBinaryTree(root)
	assert.True(t, res)
}

func TestGetLCANode(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := LowestCommonAncestor(root, root.Left, root.Right)
	assert.Equal(t, root, res)

	root = genBinaryTree(originArr, 0)
	res = LowestCommonAncestor(root, root.Left, root.Left.Left)
	assert.Equal(t, root.Left, res)

	root = genBinaryTree(originArr, 0)
	res = LowestCommonAncestor(root, root.Left.Right, root.Left.Left)
	assert.Equal(t, root.Left, res)

	root = genBinaryTree(originArr, 0)
	res = LowestCommonAncestor(root, root, root.Left.Left)
	assert.Equal(t, root, res)
}

func TestInorderSuccessor(t *testing.T) {
	root := genBinaryTree(originArr, 0)
	res := InorderSuccessor(root, root.Left)
	assert.Equal(t, root.Left.Right, res)

	root = genBinaryTree(originArr, 0)
	res = InorderSuccessor(root, root.Left.Left)
	assert.Equal(t, root.Left, res)

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
