package services

import "github.com/tanvir-005/event-explorer/models"

type DefaultEventService struct {
	ticketmaster TicketmasterService
	cache        EventCache
}

func NewEventService(
	ticketmaster TicketmasterService,
	cache EventCache,
) *DefaultEventService {
	return &DefaultEventService{
		ticketmaster: ticketmaster,
		cache:        cache,
	}
}

func (s *DefaultEventService) GetEvent(
	eventID string,
) (models.Event, error) {
	return s.ticketmaster.GetEvent(eventID)
}

func (s *DefaultEventService) GetEvents(
	city string,
	countryCode string,
) (
	music []models.Event,
	sports []models.Event,
	musicErr error,
	sportsErr error,
) {
	music = s.getCategory(
		city,
		countryCode,
		models.MusicCategory,
		&musicErr,
	)

	sports = s.getCategory(
		city,
		countryCode,
		models.SportsCategory,
		&sportsErr,
	)

	return
}

func (s *DefaultEventService) getCategory(
	city string,
	countryCode string,
	category models.EventCategory,
	errOut *error,
) []models.Event {
	key := models.EventCacheKey{
		City:        city,
		CountryCode: countryCode,
		Category:    category,
	}

	if events, ok := s.cache.Get(key); ok {
		return events
	}

	events, err := s.ticketmaster.GetEvents(
		city,
		countryCode,
		category,
	)

	if err != nil {
		*errOut = err
		return nil
	}

	s.cache.Set(key, events)

	return events
}
