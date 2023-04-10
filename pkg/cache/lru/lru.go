package lru

import (
	"sync"

	"github.com/prichrd/concepts/pkg/container/list"
)

// Cache impelements a Least Recent Used (LRU) cache holding a
// limited set of elements. Once the capacity is reached, the
// least recently used element will be evicted from the cache,
// and the memory will be freed by garbage collection.
type Cache[K comparable, V any] struct {
	capacity int
	m        map[K]*list.Element[*cacheTuple[K, V]]
	l        *list.List[*cacheTuple[K, V]]
	mtx      sync.RWMutex
}

// cacheTuple reresents a key-value pair. It is the unit stored
// in the doubly linked list that backs the cache. That tuple
// structure allows us to get the map index from a linked list
// element, which is useful when removing the least frequently
// accessed element.
type cacheTuple[K comparable, V any] struct {
	k K
	v V
}

// New returns a configured instance of Cache.
func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		m:        make(map[K]*list.Element[*cacheTuple[K, V]]),
		l:        list.NewList[*cacheTuple[K, V]](),
	}
}

// Get returns the cached element for a key. In cases where the
// element is not present or got evicted from the cache, it returns
// a nil pointer. Because this is an LRU cache, a node is brought to
// the top of the list on every cache hit. The last element of the
// list is the first one to get evicted if a new element gets added.
func (c *Cache[K, V]) Get(k K) *V {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	el, exists := c.m[k]
	if !exists {
		return nil
	}

	c.l.MoveToFront(el)
	return &el.Val().v
}

// Put adds an element to the cache under the given key. If the cache
// capacity is reached, the least used element (last element of the
// doubly linked list) is removed.
func (c *Cache[K, V]) Put(k K, v V) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// In cases where the key already exists, we update the value.
	if _, exists := c.m[k]; exists {
		n := c.m[k]
		n.Val().v = v
		c.l.MoveToFront(n)
		return
	}

	// If the capacity is exceeded with the new element, we remove the
	// last element to make room for the new one.
	if c.l.Len()+1 > c.capacity {
		removed := c.l.Back()
		delete(c.m, removed.Val().k)
		c.l.Remove(removed)
	}

	n := &cacheTuple[K, V]{
		k: k,
		v: v,
	}
	el := c.l.PushFront(n)
	c.m[k] = el
}
