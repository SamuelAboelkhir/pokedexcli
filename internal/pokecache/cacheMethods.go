package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Before(time.Now().UTC().Add(-c.interval)) {
			delete(c.cache, key)
		}
	}
}
