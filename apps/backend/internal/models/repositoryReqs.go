package models

import "time"

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
