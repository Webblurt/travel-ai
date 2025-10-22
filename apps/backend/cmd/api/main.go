package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API endpoint
	mux.HandleFunc("/api/itinerary", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"AI itinerary generated"}`))
	})

	// Serve static frontend files
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
