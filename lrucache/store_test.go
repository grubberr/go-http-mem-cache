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
