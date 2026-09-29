package services

import (
	"log"
	"sync"

	"github.com/tanvir-005/event-explorer/models"
)

type MemoryEventCache struct {
	mu      sync.RWMutex
	entries map[models.EventCacheKey]models.EventCacheEntry
}

func NewMemoryEventCache() *MemoryEventCache {
	return &MemoryEventCache{
		entries: make(map[models.EventCacheKey]models.EventCacheEntry),
	}
}

func (c *MemoryEventCache) Get(
	key models.EventCacheKey,
) ([]models.Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	log.Printf(
		"cache hit: city=%s country=%s category=%s",
		key.City,
		key.CountryCode,
		key.Category,
	)

	return entry.Events, true
}

func (c *MemoryEventCache) Set(
	key models.EventCacheKey,
	events []models.Event,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = models.EventCacheEntry{
		Events: events,
	}
}

func (c *MemoryEventCache) Invalidate(
	key models.EventCacheKey,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.entries, key)
}

func (c *MemoryEventCache) InvalidateCity(
	city string,
	countryCode string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key := range c.entries {
		if key.City == city && key.CountryCode == countryCode {
			delete(c.entries, key)
		}
	}
}

func (c *MemoryEventCache) InvalidateAll() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[models.EventCacheKey]models.EventCacheEntry)
}
