package models

type GetItinerary struct {
	ID            string
	City          string
	ItineraryCost string
	Days          []GetDayPlan
}

type GetDayPlan struct {
	DayNumber int
	Morning   GetItemPlan
	Afternoon GetItemPlan
	Evening   GetItemPlan
}

type GetItemPlan struct {
	ID           string
	City         string
	Place        string
	Description  string
	ActivityType string
	Cost         string
	Duration     string
	Address      string
	Details      struct {
		Images []struct {
			URL         string
			Description string
		}

		Overview struct {
			ShortDescription    string
			HistoryOrBackground string
			InterestingFacts    []string
		}

		VisitingInfo struct {
			OpeningHours         string
			BestTimeToVisit      string
			AverageVisitDuration string
			EntryFee             string
			Website              string
			ContactPhone         string
		}

		LocationInfo struct {
			Neighborhood           string
			NearestTransport       string
			DistanceFromCityCenter string
			MapLink                string
		}

		Tips struct {
			WhatToBring []string
			LocalTips   []string
			SafetyNotes []string
		}
	}
}
