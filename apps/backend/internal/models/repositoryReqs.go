package models

import "time"

type ItinerariesFilters struct {
	Limit  int
	Offset int
	City   string
	Budget string
	UserID string
}

type ItineraryFilters struct {
	ID        string
	DayNum    string
	TimeOfDay string
}

type Entity struct {
	EntityName        string // must match the table name
	StringParameters  map[string]string
	IntegerParameters map[string]int
	TimeParameters    map[string]time.Time
}
