package lrucache

import (
	"slices"
	"testing"
)

func TestLRUCacheGetNotFound(t *testing.T) {
	cache := NewLRUCache(1)
	value, ok := cache.Get("key")
	if ok != false {
		t.Errorf("got %v; want %v", ok, false)
	}
	if value != "" {
		t.Errorf("got %s; want empty string", value)
	}
}

func TestLRUCacheGetFound(t *testing.T) {
	cache := NewLRUCache(1)
	cache.cache["key"] = &Item{key: "key", value: "value", element: cache.queue.PushFront("key")}

	value, ok := cache.Get("key")
	if ok != true {
		t.Errorf("got %v; want %v", ok, true)
	}
	if value != "value" {
		t.Errorf("got %s; want 'value'", value)
	}
}

func TestLRUCacheSetNew(t *testing.T) {
	cache := NewLRUCache(1)
	cache.Set("key", "value")
	item, ok := cache.cache["key"]
	if ok != true {
		t.Errorf("got %v; want %v", ok, true)
	}
	if item.key != "key" {
		t.Errorf("got %s; want 'key'", item.key)
	}
	if item.value != "value" {
		t.Errorf("got %s; want 'value'", item.value)
	}
	if queueLen := cache.queue.Len(); queueLen != 1 {
		t.Errorf("got %d; want 1", queueLen)
	}
	if queueValue := cache.queue.Back().Value; queueValue != "key" {
		t.Errorf("got %s; want 'key'", queueValue)
	}
}

func TestLRUCacheSetUpdate(t *testing.T) {
	cache := NewLRUCache(2)
	cache.cache["key1"] = &Item{key: "key1", value: "value1", element: cache.queue.PushFront("key1")}
	cache.cache["key2"] = &Item{key: "key2", value: "value2", element: cache.queue.PushFront("key2")}
	cache.cache["key3"] = &Item{key: "key3", value: "value3", element: cache.queue.PushFront("key3")}

	cache.Set("key1", "value1-update")

	if lenCache := len(cache.cache); lenCache != 3 {
		t.Errorf("got %d; want 3", lenCache)
	}

	item, ok := cache.cache["key1"]
	if ok != true {
		t.Errorf("got %v; want %v", ok, true)
	}
	if item.key != "key1" {
		t.Errorf("got %s; want 'key1'", item.key)
	}
	if item.value != "value1-update" {
		t.Errorf("got %s; want 'value1-update'", item.value)
	}

	keys := make([]string, 0, cache.queue.Len())
	for e := cache.queue.Back(); e != nil; e = e.Prev() {
		keys = append(keys, e.Value.(string))
	}

	if !slices.Equal(keys, []string{"key2", "key3", "key1"}) {
		t.Errorf("got %s; want [key2 key3 key1]", keys)
	}
}
