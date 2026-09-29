package models

type EventCategory string

const (
	MusicCategory  EventCategory = "Music"
	SportsCategory EventCategory = "Sports"
)

type Event struct {
	ID          string
	Name        string
	ImageURL    string
	Date        string
	Venue       string
	City        string
	CountryCode string
	Description string
	TicketURL   string
	Category    EventCategory
}
