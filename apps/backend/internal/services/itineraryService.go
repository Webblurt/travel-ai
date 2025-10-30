package services

import (
	"fmt"
	"time"
	cl "travel-ai/internal/clients"
	models "travel-ai/internal/models"
)

func (s *Service) GetItinerary(filters models.GetItineraryReq) (models.GetItineraryResp, error) {
	s.log.Debug("==== [GetItinerary] Started fetching itinerary ====")

	s.log.Debug("[GetItinerary] Selecting ItineraryGenerator client...")
	client, err := s.SelectClientByName(s.components.ItineraryGenerator)
	if err != nil {
		s.log.Error("Error while selecting client: ", err)
		return models.GetItineraryResp{}, err
	}
	s.log.Debug("[GetItinerary] Client selected successfully")

	reqToClient := models.OpenAIGenerateItineraryReq{
		City:     filters.City,
		Days:     filters.Days,
		Budget:   filters.Budget,
		Currency: filters.Currency,
	}

	clientResp, err := cl.AskOpenAI[models.OpenAIGenerateItineraryResp](client, "itinerary-generation", reqToClient)
	if err != nil {
		s.log.Error("Client call failed: ", err)
		return models.GetItineraryResp{}, err
	}

	itineraryResp := clientResp

	s.log.Debug("[GetItinerary] Saving response to db...")
	_, err = s.createItinerary(*itineraryResp, filters.UserID)
	if err != nil {
		s.log.Error("DB call failed: ", err)
		return models.GetItineraryResp{}, err
	}

	s.log.Debug("[GetItinerary] Mapping response records into internal models...")
	itinerary := models.GetItineraryResp{
		City:          itineraryResp.City,
		ItineraryCost: itineraryResp.ItineraryCost,
		Days:          itineraryResp.Days,
	}

	return itinerary, nil
}

func (s *Service) GetExactItinerary(id string) (models.ExactIniterary, error) {
	s.log.Debug("==== [GetExactItinerary] Started fetching itinerary ====")

	s.log.Debug("[GetExactItinerary] Preparing filters for repository")
	reqToRepo := models.ItineraryFilters{
		ID: id,
	}
	s.log.Debug("[GetExactItinerary] Selecting from db...")
	repoResp, err := s.repository.GetItinerary(s.ctx, reqToRepo)
	if err != nil {
		s.log.Error("DB call failed: ", err)
		return models.ExactIniterary{}, err
	}

	s.log.Debug("[GetExactItinerary] Mapping response records into internal models...")
	switch v := repoResp.(type) {
	case models.GetItinerary:
		return models.ConvertGetItineraryToExact(v), nil
	default:
		return models.ExactIniterary{}, fmt.Errorf("unexpected type returned from repository: %T", v)
	}
}

func (s *Service) GetExactDayItinerary(id, dayNum string) (models.ExactDayPlan, error) {
	s.log.Debug("==== [GetExactDayItinerary] Started fetching itinerary exact day ====")

	s.log.Debug("[GetExactDayItinerary] Preparing filters for repository")
	reqToRepo := models.ItineraryFilters{
		ID:     id,
		DayNum: dayNum,
	}
	s.log.Debug("[GetExactDayItinerary] Selecting from db...")
	repoResp, err := s.repository.GetItinerary(s.ctx, reqToRepo)
	if err != nil {
		s.log.Error("DB call failed: ", err)
		return models.ExactDayPlan{}, err
	}

	s.log.Debug("[GetExactDayItinerary] Mapping response records into internal models...")
	switch v := repoResp.(type) {
	case models.GetDayPlan:
		return models.ConvertGetDayToExact(v), nil
	default:
		return models.ExactDayPlan{}, fmt.Errorf("unexpected type returned from repository: %T", v)
	}
}

func (s *Service) GetTimeOfDayItinerary(id, dayNum, timeOfDay, userID string) (models.ExactItemPlan, error) {
	s.log.Debug("==== [GetTimeOfDayItinerary] Started fetching itinerary exact day ====")

	s.log.Debug("[GetTimeOfDayItinerary] Preparing filters for repository")
	reqToRepo := models.ItineraryFilters{
		ID:        id,
		DayNum:    dayNum,
		TimeOfDay: timeOfDay,
	}
	s.log.Debug("[GetTimeOfDayItinerary] Selecting from db...")
	repoResp, err := s.repository.GetItinerary(s.ctx, reqToRepo)
	if err != nil {
		s.log.Error("DB call failed: ", err)
		return models.ExactItemPlan{}, err
	}

	var itemPlan models.ExactItemPlan

	s.log.Debug("[GetExactDayItinerary] Mapping response records into internal models...")
	switch v := repoResp.(type) {
	case models.GetItemPlan:
		itemPlan = models.ConvertGetItemToExact(v)
	default:
		return models.ExactItemPlan{}, fmt.Errorf("unexpected type returned from repository: %T", v)
	}

	s.log.Debug("[GetTimeOfDayItinerary] Selecting ItineraryGenerator client...")
	client, err := s.SelectClientByName(s.components.ItineraryGenerator)
	if err != nil {
		s.log.Error("Error while selecting client: ", err)
		return models.ExactItemPlan{}, err
	}
	s.log.Debug("[GetTimeOfDayItinerary] Preparing filters for client")
	reqToClient := models.OpenAIGenerateItineraryDetailsReq{
		City:         itemPlan.City,
		TimeOfDay:    timeOfDay,
		Place:        itemPlan.Place,
		Address:      itemPlan.Address,
		Description:  itemPlan.Description,
		Duration:     itemPlan.Duration,
		Cost:         itemPlan.Cost,
		ActivityType: itemPlan.ActivityType,
	}
	clientResp, err := cl.AskOpenAI[models.OpenAIGenerateItineraryDetailsResp](client, "itinerary-details-generation", reqToClient)
	if err != nil {
		s.log.Error("Client call failed: ", err)
		return models.ExactItemPlan{}, err
	}

	planResp := clientResp

	s.log.Debug("[GetTimeOfDayItinerary] Saving response to db...")
	err = s.createItineraryDetails(*planResp, userID, itemPlan.ID)
	if err != nil {
		s.log.Error("DB call failed: ", err)
		return models.ExactItemPlan{}, err
	}

	s.log.Debug("[GetTimeOfDayItinerary] Mapping response records into internal models...")
	resp := models.ExactItemPlan{
		Place:        planResp.Place,
		Description:  planResp.Description,
		ActivityType: planResp.ActivityType,
		Cost:         planResp.Cost,
		Duration:     planResp.Duration,
		Address:      planResp.Address,
		Details:      planResp.Details,
	}

	return resp, nil
}

func (s *Service) createItinerary(req models.OpenAIGenerateItineraryResp, user string) (string, error) {
	s.log.Debug("Creating db itinerary...............")
	tx, err := s.repository.BeginTx(s.ctx)
	if err != nil {
		s.log.Error("Error start transaction while itinerary creation: ", err)
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			s.log.Warn("Rollback started")
			if rollbackErr := tx.Rollback(s.ctx); rollbackErr != nil {
				s.log.Error("Rollback error: ", rollbackErr)
			}
		}
	}()

	itineraryEntity := models.Entity{
		EntityName: "itineraries",
		StringParameters: map[string]string{
			"city":       req.City,
			"cost":       req.ItineraryCost,
			"created_by": user,
		},
		IntegerParameters: make(map[string]int),
		TimeParameters:    make(map[string]time.Time),
	}

	itineraryID, err := s.repository.InsertTx(s.ctx, tx, itineraryEntity)
	if err != nil {
		s.log.Error("Error while creating itinerary: ", err)
		return "", err
	}

	for _, day := range req.Days {
		dayEntity := models.Entity{
			EntityName: "itinerary_days",
			StringParameters: map[string]string{
				"itinerary_id": itineraryID,
				"created_by":   user,
			},
			IntegerParameters: map[string]int{"day_number": day.DayNumber},
			TimeParameters:    make(map[string]time.Time),
		}

		dayID, err := s.repository.InsertTx(s.ctx, tx, dayEntity)
		if err != nil {
			return "", fmt.Errorf("insert day: %w", err)
		}

		activities := []struct {
			timeOfDay string
			data      models.ItemPlan
		}{
			{"morning", day.Morning},
			{"afternoon", day.Afternoon},
			{"evening", day.Evening},
		}

		for _, a := range activities {
			entity := models.Entity{
				EntityName: "day_activities",
				StringParameters: map[string]string{
					"day_id":               dayID,
					"time_of_day":          a.timeOfDay,
					"place":                a.data.Place,
					"activity_description": a.data.Description,
					"activity_type":        a.data.ActivityType,
					"cost":                 a.data.Cost,
					"duration":             a.data.Duration,
					"activity_address":     a.data.Address,
					"created_by":           user,
				},
				IntegerParameters: make(map[string]int),
				TimeParameters:    make(map[string]time.Time),
			}

			if _, err = s.repository.InsertTx(s.ctx, tx, entity); err != nil {
				return "", fmt.Errorf("insert activity (%s): %w", a.timeOfDay, err)
			}
		}
	}

	err = tx.Commit(s.ctx)
	if err != nil {
		s.log.Error("Error while commit transaction: ", err)
		return "", err
	}

	s.log.Debug("Itinerary saved.")
	return itineraryID, nil
}

func (s *Service) createItineraryDetails(req models.OpenAIGenerateItineraryDetailsResp, user, activityID string) (err error) {
	s.log.Debug("Creating itinerary details...")

	tx, err := s.repository.BeginTx(s.ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			s.log.Warn("Rollback started")
			if rbErr := tx.Rollback(s.ctx); rbErr != nil {
				s.log.Error("Rollback error: ", rbErr)
			}
		}
	}()

	newEntity := func(name string, params map[string]string) models.Entity {
		if params == nil {
			params = make(map[string]string)
		}
		params["created_by"] = user
		return models.Entity{
			EntityName:        name,
			StringParameters:  params,
			IntegerParameters: make(map[string]int),
			TimeParameters:    make(map[string]time.Time),
		}
	}

	detailsParams := map[string]string{
		"activity_id":               activityID,
		"short_description":         req.Details.Overview.ShortDescription,
		"history_of_background":     req.Details.Overview.HistoryOrBackground,
		"opening_hours":             req.Details.VisitingInfo.OpeningHours,
		"best_time_to_visit":        req.Details.VisitingInfo.BestTimeToVisit,
		"average_visit_duration":    req.Details.VisitingInfo.AverageVisitDuration,
		"entry_fee":                 req.Details.VisitingInfo.EntryFee,
		"website":                   req.Details.VisitingInfo.Website,
		"contact_phone":             req.Details.VisitingInfo.ContactPhone,
		"neighborhood":              req.Details.LocationInfo.Neighborhood,
		"nearest_transport":         req.Details.LocationInfo.NearestTransport,
		"distance_from_city_center": req.Details.LocationInfo.DistanceFromCityCenter,
		"map_link":                  req.Details.LocationInfo.MapLink,
	}

	detailsEntity := newEntity("activity_details", detailsParams)
	detailsID, err := s.repository.InsertTx(s.ctx, tx, detailsEntity)
	if err != nil {
		return fmt.Errorf("failed to insert activity details: %w", err)
	}

	insertItems := func(table string, items []string, field string) error {
		for _, val := range items {
			e := newEntity(table, map[string]string{
				"detail_id": detailsID,
				field:       val,
			})
			if _, err := s.repository.InsertTx(s.ctx, tx, e); err != nil {
				return fmt.Errorf("insert into %s failed: %w", table, err)
			}
		}
		return nil
	}

	for _, img := range req.Details.Images {
		e := newEntity("detail_images", map[string]string{
			"detail_id":         detailsID,
			"image_url":         img.URL,
			"image_description": img.Description,
		})
		if _, err = s.repository.InsertTx(s.ctx, tx, e); err != nil {
			return fmt.Errorf("failed to insert image: %w", err)
		}
	}

	if err = insertItems("detail_facts", req.Details.Overview.InterestingFacts, "fact"); err != nil {
		return err
	}
	if err = insertItems("detail_to_bring", req.Details.Tips.WhatToBring, "what_to_bring"); err != nil {
		return err
	}
	if err = insertItems("detail_local_tips", req.Details.Tips.LocalTips, "tip"); err != nil {
		return err
	}
	if err = insertItems("detail_safety_notes", req.Details.Tips.SafetyNotes, "note"); err != nil {
		return err
	}

	if err = tx.Commit(s.ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	s.log.Debug("Itinerary details successfully created.")
	return nil
}
