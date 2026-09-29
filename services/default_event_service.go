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

func (s *DefaultEventService) GetEvent(eventID string) (models.Event, error) {
	return s.ticketmaster.GetEvent(eventID)
}

type categoryResult struct {
	events []models.Event
	err    error
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
	musicCh := make(chan categoryResult, 1)
	sportsCh := make(chan categoryResult, 1)

	// Start Music request
	go func() {
		events, err := s.getCategory(
			city,
			countryCode,
			models.MusicCategory,
		)

		musicCh <- categoryResult{
			events: events,
			err:    err,
		}
	}()

	// Start Sports request
	go func() {
		events, err := s.getCategory(
			city,
			countryCode,
			models.SportsCategory,
		)

		sportsCh <- categoryResult{
			events: events,
			err:    err,
		}
	}()

	// Both goroutines have started before we wait for either result
	musicResult := <-musicCh
	sportsResult := <-sportsCh

	return musicResult.events, sportsResult.events, musicResult.err, sportsResult.err
}

func (s *DefaultEventService) getCategory(
	city string,
	countryCode string,
	category models.EventCategory,
) ([]models.Event, error) {
	key := models.EventCacheKey{
		City:        city,
		CountryCode: countryCode,
		Category:    category,
	}

	// hit
	if events, ok := s.cache.Get(key); ok {
		return events, nil
	}

	// miss: fetch from Ticketmaster
	events, err := s.ticketmaster.GetEvents(
		city,
		countryCode,
		category,
	)
	if err != nil {
		return nil, err
	}

	// only successful responses are cached
	s.cache.Set(key, events)

	return events, nil
}
