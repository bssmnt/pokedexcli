package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	stringMap map[string]cacheEntry
	mutex     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		stringMap: make(map[string]cacheEntry),
		mutex:     sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(k string, v []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newCache := cacheEntry{createdAt: time.Now(), val: v}
	c.stringMap[k] = newCache
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.stringMap[k]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for k, v := range c.stringMap {
				if v.createdAt.Add(interval).Before(time.Now()) {
					delete(c.stringMap, k)
				}
			}
			c.mutex.Unlock()
		}
	}
}
