package linked_list

/*
 * origin: https://leetcode.cn/problems/swap-nodes-in-pairs/
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
 * 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换)
 */

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{
		Val:-1,
		Next:head,
	}

	res := pre

	for ;head != nil && head.Next != nil; {
		tmp := head
		head = head.Next
		tmp.Next = head.Next
		head.Next = tmp
		pre.Next = head
		pre = tmp
		head = tmp.Next
	}

	return res.Next
}

