package linked_list

/*
 * origin: https://leetcode.cn/problems/palindrome-linked-list/
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 */

func IsPalindrome(head *ListNode) bool {
	// 反转前一半链表
	fast := head
	var pre *ListNode
	res := true

	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	headRight := head
	if fast!=nil {
		headRight = headRight.Next
	}

	headLeft := pre

	for ; headLeft != nil && headRight != nil; {
		if headLeft.Val != headRight.Val {
			res = false
			break
		}
		headLeft = headLeft.Next
		headRight = headRight.Next
	}
	for ;pre != nil; {
		tmp := pre.Next
		pre.Next = head
		head = pre
		pre = tmp
	}

	return res
}
