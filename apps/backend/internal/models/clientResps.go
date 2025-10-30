package models

type OpenAIGenerateItineraryResp struct {
	City          string    `json:"city"`
	ItineraryCost string    `json:"itinerary_cost"`
	Days          []DayPlan `json:"days"`
}

type OpenAIGenerateItineraryDetailsResp struct {
	Place        string `json:"place"`
	Description  string `json:"description"`
	ActivityType string `json:"activity_type"` // "sight", "cafe", "museum", "walk"
	Cost         string `json:"cost"`
	Duration     string `json:"duration"` // time to take
	Address      string `json:"address"`
	Details      struct {
		Images []struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		} `json:"images"`

		Overview struct {
			ShortDescription    string   `json:"short_description"`
			HistoryOrBackground string   `json:"history_or_background"`
			InterestingFacts    []string `json:"interesting_facts"`
		} `json:"overview"`

		VisitingInfo struct {
			OpeningHours         string `json:"opening_hours"`
			BestTimeToVisit      string `json:"best_time_to_visit"`
			AverageVisitDuration string `json:"average_visit_duration"`
			EntryFee             string `json:"entry_fee"`
			Website              string `json:"website"`
			ContactPhone         string `json:"contact_phone"`
		} `json:"visiting_info"`

		LocationInfo struct {
			Neighborhood           string `json:"neighborhood"`
			NearestTransport       string `json:"nearest_transport"`
			DistanceFromCityCenter string `json:"distance_from_city_center"`
			MapLink                string `json:"map_link"`
		} `json:"location_info"`

		Tips struct {
			WhatToBring []string `json:"what_to_bring"`
			LocalTips   []string `json:"local_tips"`
			SafetyNotes []string `json:"safety_notes"`
		} `json:"tips"`
	} `json:"details"`
}
