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
