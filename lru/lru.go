package leetgo

// https://leetcode-cn.com/problems/lru-cache-lcci/solution/goshuang-xiang-lian-biao-map-shi-xian-lru-by-pengt/

/*LinkNode for Doubly Linked List*/
type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

/*LRUCache implements lru */
type LRUCache struct {
	m          map[int]*LinkNode
	capacity   int
	head, tail *LinkNode
}

/*Constructor new */
func Constructor(capacity int) LRUCache {
	head := &LinkNode{-1, -1, nil, nil}
	tail := &LinkNode{-1, -1, nil, nil}
	head.next = tail
	tail.pre = head

	return LRUCache{}
}

/*Get data from lru*/
func (t *LRUCache) Get(key int) int {
	return 3
}

/*Put data into lru if over capacity delete least recent used */
func (t *LRUCache) Put(key int, value int) {

}
