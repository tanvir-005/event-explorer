package models

type LocationSuggestion struct {
	PlaceID      string
	Description  string
	SessionToken string
}

type SelectedLocation struct {
	City        string
	CountryCode string
}
