package api

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	interval time.Duration
	caches   map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		caches:   map[string]cacheEntry{},
	}
	go cache.cleanCache()
	return cache
}

func (c *Cache) add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.caches[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
}

func (c *Cache) get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if cacheEntry, ok := c.caches[key]; ok {
		return cacheEntry.data, true
	}

	return nil, false
}

func (c *Cache) cleanCache() {
	ticker := time.NewTicker((c.interval))
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for key, entry := range c.caches {
			if time.Since(entry.createdAt) >= c.interval {
				delete(c.caches, key)
			}
		}

		c.mu.Unlock()
	}
}
