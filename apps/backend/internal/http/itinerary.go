package http

import (
	"context"
	"encoding/json"
	"net/http"
	"travel-ai/backend/internal/ai"
)

func HandleGenerateItinerary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"method not allowed"}`))
		return
	}

	type Req struct {
		Destination string `json:"destination"`
		Days        int    `json:"days"`
	}
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	client := ai.NewAIClient()
	result, err := client.GenerateItinerary(context.Background(), req.Destination, req.Days)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
