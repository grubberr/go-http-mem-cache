package main

import (
	"fmt"
	"github.com/grubberr/go-http-mem-cache/lrucache"
)

func main() {
	cache := lrucache.NewLRUCache(5)
	cache.Set("key", "value")
	cache.Set("key2", "value2")
	fmt.Println(cache.Get("key"))
	fmt.Println(cache.Get("key"))
	fmt.Println(cache.Get("key"))
	fmt.Println(cache.Get("key"))
	fmt.Println(cache.Get("key2"))
	cache.Set("key3", "value3")
	cache.Set("key4", "value4")
	cache.Set("key5", "value5")
	cache.Set("key6", "value6")

	fmt.Println("Cache")
	cache.PrintCache()
}
