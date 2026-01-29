package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	stored map[string]cacheEntry
	mu     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.stored[key]
	return entry.val, ok
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stored[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.stored {
			if (-entry.createdAt) > interval {
				// Take current time, subtract with the time created, and see if the
				// time has passed the allowed interval
				delete(c.stored, key)
			}
		}
		c.mu.Unlock()
	}
}
