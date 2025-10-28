package itinerariescontrollers

import (
	"encoding/json"
	"net/http"
	"strconv"
	middleware "travel-ai/internal/api/middleware"
	models "travel-ai/internal/models"
	utils "travel-ai/internal/utils"
)

func (ic *ItinerariesController) GetItineraryHandler(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.FetchAuthorizationHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	_, err = ic.Service.Validate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	queryParams := r.URL.Query()
	filters := models.GetItineraryReq{
		City:     queryParams.Get("city"),
		Days:     utils.GetDefaultQueryValue(queryParams, "days", "3"),
		Budget:   utils.GetDefaultQueryValue(queryParams, "budget", "any"),
		Currency: utils.GetDefaultQueryValue(queryParams, "currency", "$"),
	}

	days, err := strconv.Atoi(filters.Days)
	if err != nil {
		http.Error(w, "invalid days parameter", http.StatusBadRequest)
		return
	}
	if days > 3 {
		http.Error(w, "More than 3 days you can have after subscription", http.StatusBadRequest)
		return
	}

	itinerary, err := ic.Service.GetItinerary(filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itinerary)
}
