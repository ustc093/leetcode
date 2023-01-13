package hot_100

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
// 返回删除后的链表的头节点。
//
// 注意：此题对比原题有改动
//
// 示例 1:
//
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//
//
// 示例 2:
//
// 输入: head = [4,5,1,9], val = 1
// 输出: [4,5,9]
// 解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{
		Next: head,
	}
	index := 0
	tmp := head
	for ; tmp != nil; tmp = tmp.Next {
		if tmp.Val == val {
			break
		}
		index++
	}
	curr := res
	for i := 0; i < index; i++ {
		curr.Next = head
		curr = head
		head = head.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	} else {
		curr.Next = nil
	}
	return res.Next
}
