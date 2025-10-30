package repositories

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	models "travel-ai/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetItinerary(ctx context.Context, filters models.ItineraryFilters) (interface{}, error) {
	r.log.Debug("Filters in repo layer: ", filters)

	query := `
		SELECT
			i.id, i.city, i.cost, da.id AS activity_id,
			id.day_number, da.time_of_day, da.place,
			da.activity_description, da.activity_type, da.cost AS activity_cost,
			da.duration, da.activity_address,
			ad.short_description, ad.history_of_background,
			df.fact, di.image_url, di.image_description,
			ad.opening_hours, ad.best_time_to_visit, ad.average_visit_duration,
			ad.entry_fee, ad.website, ad.contact_phone,
			ad.neighborhood, ad.nearest_transport, ad.distance_from_city_center, ad.map_link,
			dtb.what_to_bring, dlt.tip, dsn.note
		FROM itineraries i
		LEFT JOIN itinerary_days id ON id.itinerary_id = i.id AND id.deleted_at IS NULL
		LEFT JOIN day_activities da ON da.day_id = id.id AND da.deleted_at IS NULL
		LEFT JOIN activity_details ad ON ad.activity_id = da.id AND ad.deleted_at IS NULL
		LEFT JOIN detail_images di ON di.detail_id = ad.id AND di.deleted_at IS NULL
		LEFT JOIN detail_facts df ON df.detail_id = ad.id AND df.deleted_at IS NULL
		LEFT JOIN detail_to_bring dtb ON dtb.detail_id = ad.id AND dtb.deleted_at IS NULL
		LEFT JOIN detail_local_tips dlt ON dlt.detail_id = ad.id AND dlt.deleted_at IS NULL
		LEFT JOIN detail_safety_notes dsn ON dsn.detail_id = ad.id AND dsn.deleted_at IS NULL
		WHERE i.deleted_at IS NULL
	`

	args := []interface{}{}
	argID := 1

	if filters.ID != "" {
		query += fmt.Sprintf(" AND i.id = $%d", argID)
		args = append(args, filters.ID)
		argID++
	}
	if filters.DayNum != "" {
		daynum, err := strconv.Atoi(filters.DayNum)
		if err != nil {
			r.log.Warn("Invalid day number: ", filters.DayNum)
		}
		query += fmt.Sprintf(" AND id.day_number = $%d", argID)
		args = append(args, daynum)
		argID++
	}
	if filters.TimeOfDay != "" {
		query += fmt.Sprintf(" AND da.time_of_day ILIKE $%d", argID)
		args = append(args, "%"+filters.TimeOfDay+"%")
		argID++
	}

	query += " ORDER BY id.day_number, da.time_of_day"

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	scanItem := func(rows pgx.Rows) (models.GetItemPlan, int, string, error) {
		var item models.GetItemPlan
		var dayNum int
		var timeOfDay string
		var (
			fact, bring, tip, note *string
			imgURL, imgDesc        *string
			activityID             string
		)

		err := rows.Scan(
			new(interface{}), &item.City, new(interface{}), &activityID,
			&dayNum, &timeOfDay, &item.Place,
			&item.Description, &item.ActivityType, &item.Cost,
			&item.Duration, &item.Address,
			&item.Details.Overview.ShortDescription, &item.Details.Overview.HistoryOrBackground,
			&fact, &imgURL, &imgDesc,
			&item.Details.VisitingInfo.OpeningHours, &item.Details.VisitingInfo.BestTimeToVisit,
			&item.Details.VisitingInfo.AverageVisitDuration, &item.Details.VisitingInfo.EntryFee,
			&item.Details.VisitingInfo.Website, &item.Details.VisitingInfo.ContactPhone,
			&item.Details.LocationInfo.Neighborhood, &item.Details.LocationInfo.NearestTransport,
			&item.Details.LocationInfo.DistanceFromCityCenter, &item.Details.LocationInfo.MapLink,
			&bring, &tip, &note,
		)
		if err != nil {
			return item, 0, "", err
		}

		item.ID = activityID

		if imgURL != nil {
			item.Details.Images = append(item.Details.Images, struct {
				URL         string
				Description *string
			}{URL: *imgURL, Description: imgDesc})
		}
		if fact != nil {
			item.Details.Overview.InterestingFacts = append(item.Details.Overview.InterestingFacts, fact)
		}
		if bring != nil {
			item.Details.Tips.WhatToBring = append(item.Details.Tips.WhatToBring, bring)
		}
		if tip != nil {
			item.Details.Tips.LocalTips = append(item.Details.Tips.LocalTips, tip)
		}
		if note != nil {
			item.Details.Tips.SafetyNotes = append(item.Details.Tips.SafetyNotes, note)
		}

		return item, dayNum, timeOfDay, nil
	}

	switch {
	case filters.TimeOfDay != "":
		var result models.GetItemPlan
		for rows.Next() {
			item, _, _, err := scanItem(rows)
			if err != nil {
				return nil, err
			}
			result.ID = item.ID
			result.City = item.City
			result.Place = item.Place
			result.Description = item.Description
			result.Details.Images = append(result.Details.Images, item.Details.Images...)
			result.Details.Overview.InterestingFacts = append(result.Details.Overview.InterestingFacts, item.Details.Overview.InterestingFacts...)
			result.Details.Tips.WhatToBring = append(result.Details.Tips.WhatToBring, item.Details.Tips.WhatToBring...)
			result.Details.Tips.LocalTips = append(result.Details.Tips.LocalTips, item.Details.Tips.LocalTips...)
			result.Details.Tips.SafetyNotes = append(result.Details.Tips.SafetyNotes, item.Details.Tips.SafetyNotes...)
		}
		return result, rows.Err()

	case filters.DayNum != "":
		var day models.GetDayPlan
		for rows.Next() {
			item, dayNum, timeOfDay, err := scanItem(rows)
			if err != nil {
				return nil, err
			}
			day.DayNumber = &dayNum
			switch strings.ToLower(timeOfDay) {
			case "morning":
				day.Morning = item
			case "afternoon":
				day.Afternoon = item
			case "evening":
				day.Evening = item
			}
		}
		return day, rows.Err()

	case filters.ID != "":
		var it models.GetItinerary
		for rows.Next() {
			item, dayNum, timeOfDay, err := scanItem(rows)
			if err != nil {
				return nil, err
			}
			it.ID = filters.ID
			it.City = item.City

			var day *models.GetDayPlan
			for i := range it.Days {
				if it.Days[i].DayNumber != nil && *it.Days[i].DayNumber == dayNum {
					day = &it.Days[i]
					break
				}
			}
			if day == nil {
				it.Days = append(it.Days, models.GetDayPlan{DayNumber: &dayNum})
				day = &it.Days[len(it.Days)-1]
			}

			switch strings.ToLower(timeOfDay) {
			case "morning":
				day.Morning = item
			case "afternoon":
				day.Afternoon = item
			case "evening":
				day.Evening = item
			}
		}
		return it, rows.Err()

	default:
		return nil, fmt.Errorf("no filters provided")
	}
}

func (r *Repository) GetItinerariesCount(ctx context.Context, filters models.ItinerariesFilters) (int, error) {
	r.log.Debug("Filters in repo layer: ", filters)

	query := `
		SELECT COUNT(*)
		FROM itineraries
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argID := 1

	if filters.City != "" {
		query += fmt.Sprintf(" AND city ILIKE $%d", argID)
		args = append(args, "%"+filters.City+"%")
		argID++
	}
	if filters.Budget != "" {
		query += fmt.Sprintf(" AND cost ILIKE $%d", argID)
		args = append(args, "%"+filters.Budget+"%")
		argID++
	}
	if filters.UserID != "" {
		query += fmt.Sprintf(" AND created_by = $%d", argID)
		args = append(args, filters.UserID)
		argID++
	}

	var count int
	r.log.Debug("Query execution: ", query)
	err := r.DB.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}
	r.log.Debug("Query executed.")
	r.log.Debug("Total itineraries: ", count)

	return count, nil
}

func (r *Repository) GetRepoItineraries(ctx context.Context, filters models.ItinerariesFilters) ([]models.GetItinerary, error) {
	r.log.Debug("Filters in repo layer: ", filters)

	query := `
		SELECT
			i.id, i.city, i.cost, da.id AS activity_id,
			id.day_number, da.time_of_day, da.place,
			da.activity_description, da.activity_type, da.cost AS activity_cost,
			da.duration, da.activity_address,
			ad.short_description, ad.history_of_background,
			df.fact, di.image_url, di.image_description,
			ad.opening_hours, ad.best_time_to_visit, ad.average_visit_duration,
			ad.entry_fee, ad.website, ad.contact_phone,
			ad.neighborhood, ad.nearest_transport, ad.distance_from_city_center, ad.map_link,
			dtb.what_to_bring, dlt.tip, dsn.note
		FROM itineraries i
		LEFT JOIN itinerary_days id ON id.itinerary_id = i.id AND id.deleted_at IS NULL
		LEFT JOIN day_activities da ON da.day_id = id.id AND da.deleted_at IS NULL
		LEFT JOIN activity_details ad ON ad.activity_id = da.id AND ad.deleted_at IS NULL
		LEFT JOIN detail_images di ON di.detail_id = ad.id AND di.deleted_at IS NULL
		LEFT JOIN detail_facts df ON df.detail_id = ad.id AND df.deleted_at IS NULL
		LEFT JOIN detail_to_bring dtb ON dtb.detail_id = ad.id AND dtb.deleted_at IS NULL
		LEFT JOIN detail_local_tips dlt ON dlt.detail_id = ad.id AND dlt.deleted_at IS NULL
		LEFT JOIN detail_safety_notes dsn ON dsn.detail_id = ad.id AND dsn.deleted_at IS NULL
		WHERE i.deleted_at IS NULL
	`

	args := []interface{}{}
	argID := 1

	if filters.City != "" {
		query += fmt.Sprintf(" AND i.city ILIKE $%d", argID)
		args = append(args, "%"+filters.City+"%")
		argID++
	}
	if filters.Budget != "" {
		query += fmt.Sprintf(" AND i.cost ILIKE $%d", argID)
		args = append(args, "%"+filters.Budget+"%")
		argID++
	}
	query += fmt.Sprintf(" AND i.created_by = $%d", argID)
	args = append(args, filters.UserID)
	argID++

	query += fmt.Sprintf(" ORDER BY i.id, id.day_number, da.time_of_day DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filters.Limit, filters.Offset)

	r.log.Debug("Query execution: ", query)
	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	r.log.Debug("Query executed.")

	r.log.Debug("Start scanning...")
	scanItem := func(rows pgx.Rows) (models.GetItemPlan, string, int, string, error) {
		var item models.GetItemPlan
		var itineraryID string
		var dayNum int
		var timeOfDay string
		var (
			fact, bring, tip, note *string
			imgURL, imgDesc        *string
			activityID             string
		)

		err := rows.Scan(
			&itineraryID, &item.City, new(interface{}), &activityID,
			&dayNum, &timeOfDay, &item.Place,
			&item.Description, &item.ActivityType, &item.Cost,
			&item.Duration, &item.Address,
			&item.Details.Overview.ShortDescription, &item.Details.Overview.HistoryOrBackground,
			&fact, &imgURL, &imgDesc,
			&item.Details.VisitingInfo.OpeningHours, &item.Details.VisitingInfo.BestTimeToVisit,
			&item.Details.VisitingInfo.AverageVisitDuration, &item.Details.VisitingInfo.EntryFee,
			&item.Details.VisitingInfo.Website, &item.Details.VisitingInfo.ContactPhone,
			&item.Details.LocationInfo.Neighborhood, &item.Details.LocationInfo.NearestTransport,
			&item.Details.LocationInfo.DistanceFromCityCenter, &item.Details.LocationInfo.MapLink,
			&bring, &tip, &note,
		)
		if err != nil {
			return item, "", 0, "", err
		}

		item.ID = activityID

		if imgURL != nil {
			item.Details.Images = append(item.Details.Images, struct {
				URL         string
				Description *string
			}{URL: *imgURL, Description: imgDesc})
		}
		if fact != nil {
			item.Details.Overview.InterestingFacts = append(item.Details.Overview.InterestingFacts, fact)
		}
		if bring != nil {
			item.Details.Tips.WhatToBring = append(item.Details.Tips.WhatToBring, bring)
		}
		if tip != nil {
			item.Details.Tips.LocalTips = append(item.Details.Tips.LocalTips, tip)
		}
		if note != nil {
			item.Details.Tips.SafetyNotes = append(item.Details.Tips.SafetyNotes, note)
		}

		return item, itineraryID, dayNum, timeOfDay, nil
	}

	itineraries := make([]models.GetItinerary, 0)

	for rows.Next() {
		item, itineraryID, dayNum, timeOfDay, err := scanItem(rows)
		if err != nil {
			return nil, err
		}

		var itinerary *models.GetItinerary
		for i := range itineraries {
			if itineraries[i].ID == itineraryID {
				itinerary = &itineraries[i]
				break
			}
		}
		if itinerary == nil {
			itineraries = append(itineraries, models.GetItinerary{
				ID:   itineraryID,
				City: item.City,
			})
			itinerary = &itineraries[len(itineraries)-1]
		}

		var day *models.GetDayPlan
		for i := range itinerary.Days {
			if itinerary.Days[i].DayNumber != nil && *itinerary.Days[i].DayNumber == dayNum {
				day = &itinerary.Days[i]
				break
			}
		}
		if day == nil {
			itinerary.Days = append(itinerary.Days, models.GetDayPlan{DayNumber: &dayNum})
			day = &itinerary.Days[len(itinerary.Days)-1]
		}

		switch strings.ToLower(timeOfDay) {
		case "morning":
			day.Morning = item
		case "afternoon":
			day.Afternoon = item
		case "evening":
			day.Evening = item
		}
	}
	r.log.Debug("Scan finished.")

	return itineraries, rows.Err()
}
