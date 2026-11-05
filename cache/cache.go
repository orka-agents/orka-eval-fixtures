// Package cache is a size-bounded LRU cache whose entries expire.
package cache

import (
	"container/list"
	"sync"
	"time"
)

type entry struct {
	key     string
	value   any
	expires time.Time
}

// Cache holds at most Capacity entries and evicts the least recently used.
// It is safe for concurrent use.
type Cache struct {
	mu       sync.Mutex
	capacity int
	ttl      time.Duration
	now      func() time.Time
	order    *list.List
	items    map[string]*list.Element
}

// New returns a cache of capacity entries that each live for ttl.
func New(capacity int, ttl time.Duration) *Cache {
	if capacity < 1 {
		capacity = 1
	}
	return &Cache{
		capacity: capacity,
		ttl:      ttl,
		now:      time.Now,
		order:    list.New(),
		items:    make(map[string]*list.Element),
	}
}

// Get returns the value for key and marks it recently used. An expired entry
// is removed and reported as missing.
func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	element, ok := c.items[key]
	if !ok {
		return nil, false
	}
	e := element.Value.(*entry)
	if !c.now().Before(e.expires) {
		c.removeElement(element)
		return nil, false
	}
	c.order.MoveToFront(element)
	return e.value, true
}

// Set stores value under key, evicting the least recently used entry when
// the cache is full.
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expires := c.now().Add(c.ttl)
	if element, ok := c.items[key]; ok {
		e := element.Value.(*entry)
		e.value, e.expires = value, expires
		c.order.MoveToFront(element)
		return
	}
	c.items[key] = c.order.PushFront(&entry{key: key, value: value, expires: expires})
	if c.order.Len() > c.capacity {
		c.removeElement(c.order.Back())
	}
}

func (c *Cache) removeElement(element *list.Element) {
	c.order.Remove(element)
	delete(c.items, element.Value.(*entry).key)
}
