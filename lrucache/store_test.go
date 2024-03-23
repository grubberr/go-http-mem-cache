package lrucache

import (
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
	cache.cache["key"] = Item{key: "key", value: "value", element: cache.queue.PushFront("key")}

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
