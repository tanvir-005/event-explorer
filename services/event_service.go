package services

import "github.com/tanvir-005/event-explorer/models"

type EventService interface {
	GetEvents(
		city string,
		countryCode string,
	) (
		music []models.Event,
		sports []models.Event,
		musicErr error,
		sportsErr error,
	)

	GetEvent(
		eventID string,
	) (models.Event, error)
}
