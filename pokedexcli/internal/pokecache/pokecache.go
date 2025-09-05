package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		store: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c

}

func (c *Cache) Add(key string, val []byte) {
	// SafeGaurd the data
	c.mu.Lock()
	defer c.mu.Unlock()

	entery := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.store[key] = entery
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// SafeGaurd the data
	c.mu.RLock()
	defer c.mu.RUnlock()

	entery, ok := c.store[key]
	if ok {
		return entery.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	// SafeGaurd the strut
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.store {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.store, k)
		}
	}
}
