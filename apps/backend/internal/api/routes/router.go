package routes

import (
	"errors"
	"net/http"
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

	// =======================
	// FRONTEND
	// =======================
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	return mux, nil
}
