package itinerariescontrollers

import (
	"net/http"
	"strconv"
	models "travel-ai/internal/models"
	utils "travel-ai/internal/utils"
)

func (ic *ItinerariesController) GetItineraryHandler() http.HandlerFunc {
	return ic.withAuth(func(w http.ResponseWriter, r *http.Request, token string, userid string) {
		query := r.URL.Query()
		filters := models.GetItineraryReq{
			City:     query.Get("city"),
			Days:     utils.GetDefaultQueryValue(query, "days", "3"),
			Budget:   utils.GetDefaultQueryValue(query, "budget", "any"),
			Currency: utils.GetDefaultQueryValue(query, "currency", "$"),
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

		filters.UserID = userid

		itinerary, err := ic.Service.GetItinerary(filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(w, itinerary)
	})
}

func (ic *ItinerariesController) GetSavedItineraryHandler(id string) http.HandlerFunc {
	return ic.withAuth(func(w http.ResponseWriter, r *http.Request, token string, userid string) {
		itinerary, err := ic.Service.GetExactItinerary(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, itinerary)
	})
}

func (ic *ItinerariesController) GetExactDayItineraryHandler(id, dayNum string) http.HandlerFunc {
	return ic.withAuth(func(w http.ResponseWriter, r *http.Request, token string, userid string) {
		day, err := ic.Service.GetExactDayItinerary(id, dayNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, day)
	})
}

func (ic *ItinerariesController) GetTimeOfDayItineraryHandler(id, dayNum, section string) http.HandlerFunc {
	return ic.withAuth(func(w http.ResponseWriter, r *http.Request, token string, userid string) {
		part, err := ic.Service.GetTimeOfDayItinerary(id, dayNum, section, userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, part)
	})
}
