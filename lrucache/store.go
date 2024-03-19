package lrucache

import (
	"container/list"
	"fmt"
	"sync"
)

type Item struct {
	key     string
	value   string
	element *list.Element
}

type LRUCache struct {
	size  int
	m     sync.Mutex
	cache map[string]Item
	queue *list.List
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:  size,
		m:     sync.Mutex{},
		cache: make(map[string]Item),
		queue: list.New(),
	}
}

func (r *LRUCache) Get(key string) (string, bool) {
	r.m.Lock()
	defer r.m.Unlock()
	item, ok := r.cache[key]
	if !ok {
		return "", ok
	}
	r.queue.MoveToFront(item.element)
	return item.value, ok
}

func (r *LRUCache) Set(key, value string) {
	r.m.Lock()
	defer r.m.Unlock()

	item, ok := r.cache[key]
	if ok {
		item.value = value
		r.queue.MoveToFront(item.element)
		return
	}

	if len(r.cache) >= r.size {
		element := r.queue.Back()
		r.queue.Remove(element)
		delete(r.cache, element.Value.(string))
	}

	element := r.queue.PushFront(key)
	r.cache[key] = Item{key: key, value: value, element: element}
}

func (r *LRUCache) PrintCache() {
	for k, v := range r.cache {
		fmt.Println(k, v)
	}
}
