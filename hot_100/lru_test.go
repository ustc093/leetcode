package hot_100

import "testing"

//["LRUCache","get","put","get","put","put","get","get"]
//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]

// null,-1,null,-1,null,null,2,6

func TestLRUCache(t *testing.T) {
    //["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// [null,null,null,0,null,-1,null,-1,3,4]
	lru := ConstructorLRUCache(2)
	lru.Put(1,0)
	lru.Put(2,2)
	res := lru.Get(1)
	println(res)
	lru.Put(3,3)
	res = lru.Get(2)
	println(res)
	lru.Put(4,4)
	res = lru.Get(1)
	println(res)
	res = lru.Get(3)
	println(res)
	res = lru.Get(4)
	println(res)
}
