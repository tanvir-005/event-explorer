package services

import "github.com/tanvir-005/event-explorer/models"

type GooglePlacesService interface {
	Autocomplete(
		input string,
		sessionToken string,
	) ([]models.LocationSuggestion, error)

	GetLocation(
		placeID string,
		sessionToken string,
	) (models.SelectedLocation, error)
}
