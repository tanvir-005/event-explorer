package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/tanvir-005/event-explorer/models"
)

const ticketmasterBaseURL = "https://app.ticketmaster.com/discovery/v2"

type DefaultTicketmasterService struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewTicketmasterService(client *http.Client) *DefaultTicketmasterService {
	if client == nil {
		client = http.DefaultClient
	}

	return &DefaultTicketmasterService{
		client:  client,
		baseURL: ticketmasterBaseURL,
		apiKey:  os.Getenv("TICKETMASTER_API_KEY"),
	}
}

func (s *DefaultTicketmasterService) GetEvents(
	city string,
	countryCode string,
	category models.EventCategory,
) ([]models.Event, error) {
	params := url.Values{}
	params.Set("apikey", s.apiKey)
	params.Set("city", city)
	params.Set("countryCode", countryCode)
	params.Set("classificationName", string(category))
	params.Set("size", "6")

	var response ticketmasterEventsResponse

	if err := s.get("/events.json", params, &response); err != nil {
		return nil, err
	}

	events := make([]models.Event, 0, len(response.Embedded.Events))

	for _, event := range response.Embedded.Events {
		events = append(events, mapTicketmasterEvent(event, string(category)))
	}

	return events, nil
}

func (s *DefaultTicketmasterService) GetEvent(
	eventID string,
) (models.Event, error) {
	params := url.Values{}
	params.Set("apikey", s.apiKey)

	var response ticketmasterEvent

	if err := s.get(
		"/events/"+url.PathEscape(eventID)+".json",
		params,
		&response,
	); err != nil {
		return models.Event{}, err
	}

	return mapTicketmasterEvent(response, ""), nil
}

func (s *DefaultTicketmasterService) get(
	path string,
	params url.Values,
	target interface{},
) error {
	requestURL := s.baseURL + path

	if len(params) > 0 {
		requestURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("ticketmaster API returned status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode ticketmaster response: %w", err)
	}

	return nil
}
