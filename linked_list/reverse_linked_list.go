package linked_list

/*
 * origin: https://leetcode.cn/problems/reverse-linked-list/
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	for ; head != nil; {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func reverseListKII(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	count := 0
	var pre *ListNode
	last := head

	// judge length

	// reverse Kth list
	for count < k && head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
		count++
	}
	last.Next = reverseListKII(head, k)

	return pre
}
