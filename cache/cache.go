package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu           sync.Mutex
	cacheEntries map[string]cacheEntry
	Interval     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: map[string]cacheEntry{},
		mu:           sync.Mutex{},
		Interval:     interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.Interval)
		c.mu.Lock()

		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.Interval {
				delete(c.cacheEntries, key)
			}
		}

		c.mu.Unlock()
	}
}
