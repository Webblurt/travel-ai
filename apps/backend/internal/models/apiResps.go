package models

type GetItineraryResp struct {
	City          string    `json:"city"`
	ItineraryCost string    `json:"itinerary_cost"`
	Days          []DayPlan `json:"days"`
}

type DayPlan struct {
	DayNumber int      `json:"day_number"`
	Morning   ItemPlan `json:"morning"`
	Day       ItemPlan `json:"day"`
	Evening   ItemPlan `json:"evening"`
}

type ItemPlan struct {
	Place        string `json:"place"`
	Description  string `json:"description"`
	ActivityType string `json:"activity_type"` // "sight", "cafe", "museum", "walk"
	Cost         string `json:"cost"`
	Duration     string `json:"duration"` // time to take
	Address      string `json:"address"`
}
