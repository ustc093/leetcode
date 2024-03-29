package hot_100

/**
请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]

示例 4：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

链接：https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head
	originNewNodeMap := make(map[*Node]*Node)
	for ; curr != nil; curr = curr.Next {
		originNewNodeMap[curr] = &Node{Val: curr.Val}
	}

	curr = head
	for ; curr != nil; curr = curr.Next {
		originNewNodeMap[curr].Random = originNewNodeMap[curr.Random]
		originNewNodeMap[curr].Next = originNewNodeMap[curr.Next]
	}

	return originNewNodeMap[head]
}

func copyRandomListII(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head
	originNewNodeMap := make(map[*Node]*Node)
	for ; curr != nil; curr = curr.Next {
		originNewNodeMap[curr] = &Node{Val: curr.Val}
	}

	curr = head
	for ; curr != nil; curr = curr.Next {
		originNewNodeMap[curr].Random = originNewNodeMap[curr.Random]
		originNewNodeMap[curr].Next = originNewNodeMap[curr.Next]
	}

	return originNewNodeMap[head]
}
