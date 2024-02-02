package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	ttl     time.Duration
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func New(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		ttl:     ttl,
		mu:      &sync.RWMutex{},
	}
	go c.readLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if entry, ok := c.entries[key]; ok {
		return entry.value, true
	}
	return nil, false
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.ttl)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
