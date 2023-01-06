package hot_100

import "container/list"

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
[[],[3],[],[],[]]
输出：[null,null,3,-1,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]

链接：https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
*/

type CQueue struct {
	stack   *list.List
	stackII *list.List
}

func CQueueConstructor() CQueue {
	return CQueue{
		stack:   list.New(),
		stackII: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.stack.PushBack(value)
}

func (c *CQueue) DeleteHead() int {
	res := -1

	if c.stack.Len() == 0 && c.stackII.Len() == 0 {
		return res
	}

	if c.stackII.Len() == 0 {

		for ; c.stack.Len() != 0; {
			ele := c.stack.Back()
			c.stack.Remove(ele)
			c.stackII.PushBack(ele.Value)
		}

	}

	ele := c.stackII.Back()
	res = ele.Value.(int)
	c.stackII.Remove(ele)

	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
