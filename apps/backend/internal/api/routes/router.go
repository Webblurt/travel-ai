package routes

import (
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
			itineraries.GetItineraryHandler()(w, r)
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
			itineraries.GetItineraryHandler()(w, r)
		}
		parts := strings.Split(path, "/")

		switch len(parts) {
		case 1:
			// /api/v1/itineraries/:id
			id := parts[0]
			itineraries.GetSavedItineraryHandler(id)(w, r)

		case 2:
			// /api/v1/itineraries/:id/:dayNum
			id := parts[0]
			dayNum := parts[1]
			itineraries.GetExactDayItineraryHandler(id, dayNum)(w, r)

		case 3:
			// /api/v1/itineraries/:id/:dayNum/:section
			id := parts[0]
			dayNum := parts[1]
			section := parts[2]

			switch section {
			case "morning", "afternoon", "evening":
				itineraries.GetTimeOfDayItineraryHandler(id, dayNum, section)(w, r)
			default:
				http.Error(w, "unknown time of a day", http.StatusBadRequest)
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
