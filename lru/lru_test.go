package leetgo

import (
	"testing"
)

// # LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// # cache.put(1, 1);
// # cache.put(2, 2);
// # cache.get(1);       // 返回  1
// # cache.put(3, 3);    // 该操作会使得密钥 2 作废
// # cache.get(2);       // 返回 -1 (未找到)
// # cache.put(4, 4);    // 该操作会使得密钥 1 作废
// # cache.get(1);       // 返回 -1 (未找到)
// # cache.get(3);       // 返回  3
// # cache.get(4);       // 返回  4
func TestLru(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	res := cache.Get(1)
	if res != 1 {
		t.Errorf("Want 1 bug got %v", res)
	}
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	final := cache.Get(4)
	if final != 4 {
		t.Errorf("Want 4 bug got %v", final)
	}
}
