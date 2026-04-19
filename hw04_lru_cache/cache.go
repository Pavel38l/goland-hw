package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.items[key]; ok {
		val.Value = cacheItem{key: key, value: value}
		c.queue.MoveToFront(val)
		return true
	}
	if c.queue.Len() == c.capacity {
		delete(c.items, c.queue.Back().Value.(cacheItem).key)
		c.queue.Remove(c.queue.Back())
	}
	valueListItem := c.queue.PushFront(cacheItem{key: key, value: value})
	c.items[key] = valueListItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		return val.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
