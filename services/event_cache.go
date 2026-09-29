package services

import "github.com/tanvir-005/event-explorer/models"

type EventCache interface {
	Get(key models.EventCacheKey) ([]models.Event, bool)
	Set(key models.EventCacheKey, events []models.Event)
	Invalidate(key models.EventCacheKey)
	InvalidateCity(city string, countryCode string)
	InvalidateAll()
}
