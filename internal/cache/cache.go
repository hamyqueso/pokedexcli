// Package pokecache handles caching of data to speed up request service

package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()

	ce, ok := c.cache[key]

	return ce.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		c.Lock()
		for key, ce := range c.cache {
			if time.Since(ce.createdAt) > interval {
				delete(c.cache, key)
			}
		}

		c.Unlock()
	}

}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
	}

	go c.reapLoop(interval)

	return c
}
