package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	value []byte
	time  time.Time
}

func NewCache(interval int) Cache {
	cachedItems := Cache{
		items: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}

	go cachedItems.ReadLoop(interval)

	return cachedItems
}

func (c *Cache) Get(api string) ([]byte, error) {

	item, ok := c.items[api]

	if !ok {
		return item.value, errors.New("404")
	}

	return item.value, nil
}

func (c *Cache) Add(api string, value []byte) error {
	c.mu.Lock()
	c.items[api] = cacheEntry{
		value: value,
		time:  time.Now(),
	}
	c.mu.Unlock()

	return nil
}

func (c *Cache) ReadLoop(interval int) {
	ticker := time.NewTicker(time.Second + 1)
	for tick := range ticker.C {
		if tick.Minute()%interval == 0 {
			c.UpdateCache(interval)
		}
	}

}

func (c *Cache) UpdateCache(interval int) error {
	currentTime := time.Now()

	for api, item := range c.items {
		changeInTime := item.time.Sub(currentTime)
		if changeInTime.Minutes() < float64(interval) {
			c.mu.Lock()
			delete(c.items, api)
			c.mu.Unlock()
		}
	}

	return nil
}
