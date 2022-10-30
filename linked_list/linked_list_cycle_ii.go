package linked_list

/*
 * origin: https://leetcode.cn/problems/linked-list-cycle-ii/
 * 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 
 * 如果链表无环，则返回 null。
 *
 * 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
 * 不允许修改 链表。
 */

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var res *ListNode
	fast := head
	slow := head
	isCycle := false
	for ;fast.Next != nil && fast.Next.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			isCycle = true
			break;
		}
	}

	if isCycle {
		tmp := head
		for ;tmp != slow; {
			tmp = tmp.Next
			slow = slow.Next
		}
		res = tmp
	}

	return res
}