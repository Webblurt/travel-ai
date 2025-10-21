package http

import (
	"encoding/json"
	"net/http"
)

type ItineraryResponse struct {
	Days []struct {
		Day  int      `json:"day"`
		Plan []string `json:"plan"`
	} `json:"days"`
}

func HandleGenerateItinerary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"method not allowed"}`))
		return
	}

	response := map[string]interface{}{
		"days": []map[string]interface{}{
			{
				"day": 1,
				"plan": []string{
					"Arrive and check-in",
					"Explore city center",
					"Visit local restaurant",
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
