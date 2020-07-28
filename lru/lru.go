package leetgo

/*LinkNode for Doubly Linked List*/
type LinkNode struct {
	key, value int
	prev, next *LinkNode
}

/*LRUCache implements lru */
type LRUCache struct {
	m map[int]*LinkNode
}

/*Constructor new */
func Constructor(capacity int) LRUCache {

	return LRUCache{}
}

/*Get data from lru*/
func (t *LRUCache) Get(key int) int {
	return 3
}

/*Put data into lru if over capacity delete least recent used */
func (t *LRUCache) Put(key int, value int) {

}
