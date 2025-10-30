package models

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func safeStrings(ptrs []*string) []string {
	if len(ptrs) == 0 {
		return nil
	}
	res := make([]string, 0, len(ptrs))
	for _, p := range ptrs {
		if p != nil {
			res = append(res, *p)
		}
	}
	return res
}

func safeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func ConvertGetItemToExact(g GetItemPlan) ExactItemPlan {
	e := ExactItemPlan{
		ID:           g.ID,
		City:         safeString(g.City),
		Place:        safeString(g.Place),
		Description:  safeString(g.Description),
		ActivityType: safeString(g.ActivityType),
		Cost:         safeString(g.Cost),
		Duration:     safeString(g.Duration),
		Address:      safeString(g.Address),
	}

	// Images
	for _, img := range g.Details.Images {
		e.Details.Images = append(e.Details.Images, struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		}{
			URL:         img.URL,
			Description: safeString(img.Description),
		})
	}

	// Overview
	e.Details.Overview.ShortDescription = safeString(g.Details.Overview.ShortDescription)
	e.Details.Overview.HistoryOrBackground = safeString(g.Details.Overview.HistoryOrBackground)
	e.Details.Overview.InterestingFacts = safeStrings(g.Details.Overview.InterestingFacts)

	// VisitingInfo
	e.Details.VisitingInfo = struct {
		OpeningHours         string `json:"opening_hours"`
		BestTimeToVisit      string `json:"best_time_to_visit"`
		AverageVisitDuration string `json:"average_visit_duration"`
		EntryFee             string `json:"entry_fee"`
		Website              string `json:"website"`
		ContactPhone         string `json:"contact_phone"`
	}{
		OpeningHours:         safeString(g.Details.VisitingInfo.OpeningHours),
		BestTimeToVisit:      safeString(g.Details.VisitingInfo.BestTimeToVisit),
		AverageVisitDuration: safeString(g.Details.VisitingInfo.AverageVisitDuration),
		EntryFee:             safeString(g.Details.VisitingInfo.EntryFee),
		Website:              safeString(g.Details.VisitingInfo.Website),
		ContactPhone:         safeString(g.Details.VisitingInfo.ContactPhone),
	}

	// LocationInfo
	e.Details.LocationInfo = struct {
		Neighborhood           string `json:"neighborhood"`
		NearestTransport       string `json:"nearest_transport"`
		DistanceFromCityCenter string `json:"distance_from_city_center"`
		MapLink                string `json:"map_link"`
	}{
		Neighborhood:           safeString(g.Details.LocationInfo.Neighborhood),
		NearestTransport:       safeString(g.Details.LocationInfo.NearestTransport),
		DistanceFromCityCenter: safeString(g.Details.LocationInfo.DistanceFromCityCenter),
		MapLink:                safeString(g.Details.LocationInfo.MapLink),
	}

	// Tips
	e.Details.Tips = struct {
		WhatToBring []string `json:"what_to_bring"`
		LocalTips   []string `json:"local_tips"`
		SafetyNotes []string `json:"safety_notes"`
	}{
		WhatToBring: safeStrings(g.Details.Tips.WhatToBring),
		LocalTips:   safeStrings(g.Details.Tips.LocalTips),
		SafetyNotes: safeStrings(g.Details.Tips.SafetyNotes),
	}

	return e
}

func ConvertGetDayToExact(g GetDayPlan) ExactDayPlan {
	return ExactDayPlan{
		DayNumber: safeInt(g.DayNumber),
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
		City:          safeString(g.City),
		ItineraryCost: safeString(g.ItineraryCost),
		Days:          exactDays,
	}
}
