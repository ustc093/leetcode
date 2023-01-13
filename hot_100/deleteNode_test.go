package hot_100

import "testing"

func TestDeleteNode(t *testing.T) {
	res := deleteNode(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
		},
	}, 1)
	println(res)
}
