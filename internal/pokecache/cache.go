package pokecache

import (
	//"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:       sync.Mutex{},
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	var en cacheEntry
	en.createdAt = time.Now().Add(c.interval)
	en.val = val

	c.mu.Lock()
	c.entries[key] = en
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.entries[key]
	c.mu.Unlock()
	if exists {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {

	for key, value := range c.entries {
		if time.Now().Sub(value.createdAt) > c.interval {
			c.mu.Lock()
			delete(c.entries, key)
			c.mu.Unlock()
		}
	}
}
