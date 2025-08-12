// Package pokecache handles caching of data to speed up request service

package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex()
}

func NewCache(interval time.Duration) Cache {
	return Cache {
		createdAt: time.Time,
	}
}
