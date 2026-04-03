package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte // raw data were caching
}

type Cache struct {
	mu      *sync.Mutex
	entries map[string]CacheEntry
}

func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := CacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.entries[k] = entry
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.entries[k]
	return v.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k := range c.entries {
			if time.Now().After(c.entries[k].createdAt.Add(interval)) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:      &sync.Mutex{},
		entries: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)

	return cache
}
