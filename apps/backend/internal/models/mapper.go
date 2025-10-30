package models

func ConvertGetItemToExact(g GetItemPlan) ExactItemPlan {
	e := ExactItemPlan{
		ID:           g.ID,
		City:         g.City,
		Place:        g.Place,
		Description:  g.Description,
		ActivityType: g.ActivityType,
		Cost:         g.Cost,
		Duration:     g.Duration,
		Address:      g.Address,
	}

	for _, img := range g.Details.Images {
		e.Details.Images = append(e.Details.Images, struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		}{URL: img.URL, Description: img.Description})
	}

	e.Details.Overview.ShortDescription = g.Details.Overview.ShortDescription
	e.Details.Overview.HistoryOrBackground = g.Details.Overview.HistoryOrBackground
	e.Details.Overview.InterestingFacts = append(e.Details.Overview.InterestingFacts, g.Details.Overview.InterestingFacts...)

	e.Details.VisitingInfo = struct {
		OpeningHours         string `json:"opening_hours"`
		BestTimeToVisit      string `json:"best_time_to_visit"`
		AverageVisitDuration string `json:"average_visit_duration"`
		EntryFee             string `json:"entry_fee"`
		Website              string `json:"website"`
		ContactPhone         string `json:"contact_phone"`
	}{
		OpeningHours:         g.Details.VisitingInfo.OpeningHours,
		BestTimeToVisit:      g.Details.VisitingInfo.BestTimeToVisit,
		AverageVisitDuration: g.Details.VisitingInfo.AverageVisitDuration,
		EntryFee:             g.Details.VisitingInfo.EntryFee,
		Website:              g.Details.VisitingInfo.Website,
		ContactPhone:         g.Details.VisitingInfo.ContactPhone,
	}

	e.Details.LocationInfo = struct {
		Neighborhood           string `json:"neighborhood"`
		NearestTransport       string `json:"nearest_transport"`
		DistanceFromCityCenter string `json:"distance_from_city_center"`
		MapLink                string `json:"map_link"`
	}{
		Neighborhood:           g.Details.LocationInfo.Neighborhood,
		NearestTransport:       g.Details.LocationInfo.NearestTransport,
		DistanceFromCityCenter: g.Details.LocationInfo.DistanceFromCityCenter,
		MapLink:                g.Details.LocationInfo.MapLink,
	}

	e.Details.Tips = struct {
		WhatToBring []string `json:"what_to_bring"`
		LocalTips   []string `json:"local_tips"`
		SafetyNotes []string `json:"safety_notes"`
	}{
		WhatToBring: append([]string{}, g.Details.Tips.WhatToBring...),
		LocalTips:   append([]string{}, g.Details.Tips.LocalTips...),
		SafetyNotes: append([]string{}, g.Details.Tips.SafetyNotes...),
	}

	return e
}

func ConvertGetDayToExact(g GetDayPlan) ExactDayPlan {
	return ExactDayPlan{
		DayNumber: g.DayNumber,
		Morning:   ConvertGetItemToExact(g.Morning),
		Afternoon: ConvertGetItemToExact(g.Afternoon),
		Evening:   ConvertGetItemToExact(g.Evening),
	}
}

func ConvertGetItineraryToExact(g GetItinerary) ExactIniterary {
	var exactDays []ExactDayPlan
	for _, d := range g.Days {
		exactDays = append(exactDays, ConvertGetDayToExact(d))
	}

	return ExactIniterary{
		ID:            g.ID,
		City:          g.City,
		ItineraryCost: g.ItineraryCost,
		Days:          exactDays,
	}
}
