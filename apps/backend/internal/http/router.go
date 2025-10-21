package http

import "net/http"

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/itinerary", HandleGenerateItinerary)
}
