package main

import (
	"log"
	"net/http"
    "encoding/json"
    "fmt"
)

func main() {
	mux := http.NewServeMux()

	// API endpoint
	mux.HandleFunc("/api/itinerary", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
        return
    }

    type Request struct {
        City string `json:"city"`
    }
    var req Request
    json.NewDecoder(r.Body).Decode(&req)

    response := map[string]string{
        "itinerary": fmt.Sprintf("Ваш 3-дневный маршрут по городу %s:\n\nДень 1 — прогулка по старому городу\nДень 2 — музеи и рестораны\nДень 3 — местные парки и виды 🌇", req.City),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
})

	// Serve static frontend files
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
