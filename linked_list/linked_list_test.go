package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReorderList(t *testing.T) {
	head := genLinkedList([]int{1,2,3,4})
	ReorderList(head)
}

func TestIsPalindrome(t *testing.T) {
	head := genLinkedList([]int{1,2,3,2,1})
	isPalindrome := IsPalindrome(head)
	assert.True(t, isPalindrome)

	head = genLinkedList([]int{1})
	assert.True(t, IsPalindrome(head))

	head = genLinkedList([]int{1,2,2,1})
	assert.True(t, IsPalindrome(head))

	head = genLinkedList([]int{1,2,2,1,3})
	assert.False(t, IsPalindrome(head))
}

func genLinkedList(arr []int) *ListNode {
	res := &ListNode{}
	if len(arr) == 0 {
		return res
	}

	curr := res
	for _, value := range arr {
		tmp := &ListNode{
			Val: value,
		}
		curr.Next = tmp
		curr = curr.Next
	}

	return res.Next
}
