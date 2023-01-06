package hot_100

import "container/list"

/**

定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

链接：https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof
*/

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

func MinStackConstructor() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (m *MinStack) Push(x int) {
	m.stack.PushBack(x)
	if m.minStack.Len() == 0 || x <= m.minStack.Back().Value.(int) {
		m.minStack.PushBack(x)
	}
}

func (m *MinStack) Pop() {
	ele := m.stack.Back()
	m.stack.Remove(ele)
	if m.minStack.Len() != 0 && ele.Value.(int) == m.minStack.Back().Value.(int) {
		m.minStack.Remove(m.minStack.Back())
	}
}

func (m *MinStack) Top() int {
	return m.stack.Back().Value.(int)
}

func (m *MinStack) Min() int {
	res := -1
	if m.minStack.Len() != 0 {
		res = m.minStack.Back().Value.(int)
	}
	return res
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
