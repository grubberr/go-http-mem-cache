package lrucache

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	key   string
	value string
	time  time.Time
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
	(*item).time = time.Now()
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

	item := &Item{key: key, value: value, time: time.Now()}
	r.cache[key] = item
	heap.Push(&r.queue, item)
}

func (r *LRUCache) PrintCache() {
	for k, v := range r.cache {
		fmt.Println(k, v)
	}
}
