package hot_100

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

origin：https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=c0glesv

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {

	res := make([]int, 0)
	if head == nil {
		return res
	}
	var pre *ListNode
	for ; head != nil; {
		tmpNext := head.Next
		head.Next = pre
		pre = head
		head = tmpNext
	}

	for ; pre != nil; pre = pre.Next {
		res = append(res, pre.Val)
	}

	return res
}
