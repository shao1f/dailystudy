package main

import "github.com/coocood/freecache"

func main() {
	cache := freecache.NewCache(512)
	cache.Set([]byte("hello"), []byte("world"), 3)
}
