package models

type EventCacheKey struct {
	City        string
	CountryCode string
	Category    EventCategory
}

type EventCacheEntry struct {
	Events []Event
}