package linked_list

/*
 * origin: https://leetcode.cn/problems/reorder-list/
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：L0 → L1 → … → Ln - 1 → Ln
 * 请将其重新排列后变为：L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	curr := head
	nodeMap := make(map[int]*ListNode, 0)
	count := 0
	for ; curr != nil; {
		nodeMap[count] = curr
		curr = curr.Next
		count++
	}
	i := 0
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	res := pre
	n := len(nodeMap)

	for ; i < (n+1)/2; i++ {
		pre.Next = nodeMap[i]
		pre = pre.Next
		if i >= n {
			break
		}
		pre.Next = nodeMap[n-i-1]
		pre = pre.Next
	}
	pre.Next = nil
	head = res.Next
}

func ReorderListII(head *ListNode) {
	if head == nil {
		return
	}
	// 找到mid
	fast := head
	mid := head
	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		mid = mid.Next
	}

	// reverse
	var pre  *ListNode
	for ;mid != nil; {
		tmp := mid.Next
		mid.Next = pre
		pre = mid
		mid = tmp
	}

	// merge pre && head
	curr := head
	for ;curr.Next != nil && pre != nil; {
		currNext := curr.Next
		preNext := pre.Next
		curr.Next = pre
		pre.Next = currNext
		curr = currNext
		pre = preNext
	}

}
