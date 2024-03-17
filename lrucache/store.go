package lrucache

import (
	"container/heap"
	"fmt"
	"sync"
)

type Item struct {
	key       string
	value     string
	frequency int
}

type LRUCache struct {
	size  int
	m     sync.RWMutex
	cache map[string]*Item
	queue Heap
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:  size,
		m:     sync.RWMutex{},
		cache: make(map[string]*Item),
		queue: Heap{},
	}
}

func (r *LRUCache) Get(key string) (string, bool) {
	r.m.RLock()
	defer r.m.RUnlock()
	item, ok := r.cache[key]
	if !ok {
		return "", ok
	}
	(*item).frequency += 1
	return (*item).value, ok
}

func (r *LRUCache) Set(key, value string) {
	_, ok := r.Get(key)
	if ok {
		return
	}

	r.m.Lock()
	defer r.m.Unlock()

	if len(r.cache) >= r.size {
		item := heap.Pop(&r.queue).(*Item)
		delete(r.cache, item.key)
	}

	item := &Item{key: key, value: value, frequency: 0}
	r.cache[key] = item
	heap.Push(&r.queue, item)
}

func (r *LRUCache) PrintCache() {
	for k, v := range r.cache {
		fmt.Println(k, v)
	}
}
