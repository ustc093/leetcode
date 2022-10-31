package linked_list

/*
 * origin: https://leetcode.cn/problems/reverse-linked-list-ii/
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
 * 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := res
	next := head.Next
	count := 1
	for ; count < right ; {
		if count < left {
			head = head.Next
			pre = pre.Next
			next = next.Next
			count++
			continue
		}
		tmp := next
		next = next.Next
		tmp.Next = head
		head = tmp
		count++
	}
	pre.Next.Next = next
	pre.Next = head

	return res.Next
}

