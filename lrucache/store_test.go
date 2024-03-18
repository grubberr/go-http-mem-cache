package lrucache

import (
	"testing"
	"time"
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
	cache.cache["key"] = &Item{key: "key", value: "value"}

	before_key_access := time.Now()
	time.Sleep(1 * time.Millisecond)
	value, ok := cache.Get("key")
	time.Sleep(1 * time.Millisecond)
	after_key_access := time.Now()

	if ok != true {
		t.Errorf("got %v; want %v", ok, true)
	}
	if value != "value" {
		t.Errorf("got %s; want 'value'", value)
	}

	if !before_key_access.Before(cache.cache["key"].time) {
		t.Errorf("Time %s has to be before %s", before_key_access, cache.cache["key"].time)
	}

	if !after_key_access.After(cache.cache["key"].time) {
		t.Errorf("Time %s has to be after %s", after_key_access, cache.cache["key"].time)
	}
}
