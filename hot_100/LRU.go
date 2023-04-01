// LT_146_LruCache_test

package hot_100

import (
	"container/list"
)

//
// 请你设计并实现一个满足 
// LRU (最近最少使用) 缓存 约束的数据结构。
// 
//
// 
// 实现 
// LRUCache 类：
// 
//
// 
// 
// 
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存 
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。 
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。 
// 
// 
// 
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。 
//
// 
//
// 示例： 
//
// 
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
// 
//
// 
//
// 提示： 
//
// 
// 1 <= capacity <= 3000 
// 0 <= key <= 10000 
// 0 <= value <= 10⁵ 
// 最多调用 2 * 10⁵ 次 get 和 put 
// 
//
// Related Topics设计 | 哈希表 | 链表 | 双向链表 
//
// 👍 2599, 👎 0 
//
//
//
//

type LRUCache struct {
	Capacity    int
	KeyValueMap map[int]int
	KeyListMap  map[int]*list.Element
	List        *list.List
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		Capacity:    capacity,
		KeyValueMap: make(map[int]int),
		KeyListMap:  make(map[int]*list.Element),
		List:        list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if value, ok := l.KeyValueMap[key]; ok {
		ele := l.KeyListMap[key]
		l.List.MoveToBack(ele)
		return value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	_, ok := l.KeyValueMap[key]
	// 驱逐最远使用
	if l.List.Len() >= l.Capacity && !ok{
		l.Capacity--
		ele := l.List.Front()
		l.List.Remove(ele)
		keyInt := ele.Value.(int)

		// delete map
		delete(l.KeyValueMap, keyInt)
		delete(l.KeyListMap, keyInt)
	}

	// 之前put过
	if ok {
		ele := l.KeyListMap[key]
		l.List.MoveToBack(ele)
	} else {
		ele := l.List.PushBack(key)
		l.KeyListMap[key] = ele
	}
	l.KeyValueMap[key] = value
}
