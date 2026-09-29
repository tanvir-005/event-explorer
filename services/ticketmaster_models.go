package services

import "github.com/tanvir-005/event-explorer/models"

type ticketmasterEventsResponse struct {
	Embedded struct {
		Events []ticketmasterEvent `json:"events"`
	} `json:"_embedded"`
}

type ticketmasterEvent struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Images      []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
	Dates struct {
		Start struct {
			LocalDate string `json:"localDate"`
			LocalTime string `json:"localTime"`
		} `json:"start"`
	} `json:"dates"`
	Embedded struct {
		Venues []struct {
			Name string `json:"name"`
			City struct {
				Name string `json:"name"`
			} `json:"city"`
			Country struct {
				CountryCode string `json:"countryCode"`
			} `json:"country"`
			Address struct {
				Line1 string `json:"line1"`
			} `json:"address"`
		} `json:"venues"`
	} `json:"_embedded"`
}

func mapTicketmasterEvent(
	event ticketmasterEvent,
	category string,
) models.Event {
	var imageURL string
	if len(event.Images) > 0 {
		imageURL = event.Images[0].URL
	}

	var venue string
	var city string
	var countryCode string

	if len(event.Embedded.Venues) > 0 {
		v := event.Embedded.Venues[0]

		venue = v.Name
		city = v.City.Name
		countryCode = v.Country.CountryCode
	}

	date := event.Dates.Start.LocalDate
	if event.Dates.Start.LocalTime != "" {
		date += " " + event.Dates.Start.LocalTime
	}

	return models.Event{
		ID:          event.ID,
		Name:        event.Name,
		ImageURL:    imageURL,
		Date:        date,
		Venue:       venue,
		City:        city,
		CountryCode: countryCode,
		Description: event.Description,
		TicketURL:   event.URL,
		Category:    models.EventCategory(category),
	}
}
