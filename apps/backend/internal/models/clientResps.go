package models

type OpenAIResp struct {
	City          string    `json:"city"`
	ItineraryCost string    `json:"itinerary_cost"`
	Days          []DayPlan `json:"days"`
}
