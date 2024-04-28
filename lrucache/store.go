package lrucache

import (
	"container/list"
	"fmt"
	"sync"
)

type Item struct {
	key            string
	value          []byte
	element        *list.Element
	access_element *list.Element
}

type KeyAccess struct {
	Key    string `json:"key"`
	Access int    `json:"access"`
}

type LRUCache struct {
	size        int
	m           sync.Mutex
	cache       map[string]*Item
	queue       *list.List
	access_list *list.List
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:        size,
		m:           sync.Mutex{},
		cache:       make(map[string]*Item),
		queue:       list.New(),
		access_list: list.New(),
	}
}

func IncreaseAccess(mylist *list.List, element *list.Element) {
	element.Value.(*KeyAccess).Access += 1
	for {
		element_prev := element.Prev()
		if element_prev == nil {
			break
		}

		if element.Value.(*KeyAccess).Access < element_prev.Value.(*KeyAccess).Access {
			break
		}
		mylist.MoveBefore(element, element_prev)
	}
}

func (r *LRUCache) Get(key string) ([]byte, bool) {
	r.m.Lock()
	defer r.m.Unlock()
	item, ok := r.cache[key]
	if !ok {
		return nil, ok
	}
	r.queue.MoveToFront(item.element)
	IncreaseAccess(r.access_list, item.access_element)
	return item.value, ok
}

func (r *LRUCache) Set(key string, value []byte) {
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
		key := element.Value.(string)
		item := r.cache[key]
		r.access_list.Remove(item.access_element)
		r.queue.Remove(element)
		delete(r.cache, key)
	}

	element := r.queue.PushFront(key)
	access_element := r.access_list.PushBack(&KeyAccess{Key: key, Access: 0})
	r.cache[key] = &Item{key: key, value: value, element: element, access_element: access_element}
}

func (r *LRUCache) GetTopKeys(n int) []*KeyAccess {
	res := make([]*KeyAccess, 0, n)
	for element := r.access_list.Front(); element != nil; element = element.Next() {
		key_access := element.Value.(*KeyAccess)
		res = append(res, key_access)
		if n -= 1; n == 0 {
			break
		}
	}
	return res
}

func (r *LRUCache) PrintCache() {
	for k, v := range r.cache {
		fmt.Println(k, v)
	}
}
