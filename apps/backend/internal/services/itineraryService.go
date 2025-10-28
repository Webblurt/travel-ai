package services

import models "travel-ai/internal/models"

func (s *Service) GetItinerary(filters models.GetItineraryReq) (models.GetItineraryResp, error) {
	s.log.Debug("==== [GetItinerary] Started fetching itinerary ====")

	s.log.Debug("[GetItinerary] Selecting ItineraryGenerator client...")
	client, err := s.SelectClientByName(s.components.ItineraryGenerator)
	if err != nil {
		s.log.Error("Error while selecting client: ", err)
		return models.GetItineraryResp{}, err
	}
	s.log.Debug("[GetItinerary] Client selected successfully")

	reqToClient := models.OpenAIReq{
		City:     filters.City,
		Days:     filters.Days,
		Budget:   filters.Budget,
		Currency: filters.Currency,
	}

	clientResp, err := client.GenerateItinerary(reqToClient)
	if err != nil {
		s.log.Error("Client call failed: ", err)
		return models.GetItineraryResp{}, err
	}

	s.log.Debug("[GetItinerary] Mapping response records into internal models...")
	itinerary := models.GetItineraryResp{
		City:          clientResp.City,
		ItineraryCost: clientResp.ItineraryCost,
		Days:          clientResp.Days,
	}

	return itinerary, nil
}
