package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache:    make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		Interval: interval,
	}
	go cache.reapLoop()
	return &cache
}
