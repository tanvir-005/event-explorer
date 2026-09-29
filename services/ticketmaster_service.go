package services

import "github.com/tanvir-005/event-explorer/models"

type TicketmasterService interface {
	GetEvents(
		city string,
		countryCode string,
		category models.EventCategory,
	) ([]models.Event, error)

	GetEvent(
		eventID string,
	) (models.Event, error)
}
