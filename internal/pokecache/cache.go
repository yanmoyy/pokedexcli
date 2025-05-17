package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exist := c.cache[key]
	if exist {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			c.mu.Lock()
			now := time.Now()
			for key, entry := range c.cache {
				if now.After(entry.createdAt.Add(interval)) {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
