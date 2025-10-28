package routes

import (
	"context"
	"errors"
	"net/http"
	"strings"
	itinerariesControllers "travel-ai/internal/api/controllers/itinerariesControllers"
	services "travel-ai/internal/services"
)

func CreateRoutes(service services.ServiceInterface) (http.Handler, error) {
	if service == nil {
		return nil, errors.New("service is nil")
	}

	itineraries := itinerariesControllers.NewItinerariesController(service)

	mux := http.NewServeMux()

	// =======================
	// ITINERARIES
	// =======================
	mux.HandleFunc("/api/v1/itineraries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itineraries.GetItineraryHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// common handler for itineraries: /itineraries/:id, /itineraries/:id/:day_number/morning, /itineraries/:id/:day_number/afternoon, /itineraries/:id/:day_number/evening,
	mux.HandleFunc("/api/v1/itineraries/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/api/v1/itineraries/")
		if path == "" {
			itineraries.GetItineraryHandler(w, r)
			return
		}

		parts := strings.Split(path, "/")

		switch len(parts) {
		case 1:
			// /api/v1/itineraries/:id
			r = r.WithContext(context.WithValue(r.Context(), "id", parts[0]))
			itineraries.GetSavedItineraryHandler(w, r)

		case 3:
			id := parts[0]
			dayNum := parts[1]
			section := parts[2]

			switch section {
			case "morning": // /api/v1/itineraries/:id/:day_number/morning
				itineraries.GetMorningItineraryHandler(w, r, id, dayNum)
			case "afternoon": // /api/v1/itineraries/:id/:day_number/afternoon
				itineraries.GetAfternoonItineraryHandler(w, r, id, dayNum)
			case "evening": // /api/v1/itineraries/:id/:day_number/evening
				itineraries.GetEveningItineraryHandler(w, r, id, dayNum)
			default:
				http.Error(w, "unknown section", http.StatusBadRequest)
			}

		default:
			http.NotFound(w, r)
		}
	})

	// =======================
	// FRONTEND
	// =======================
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	return mux, nil
}
