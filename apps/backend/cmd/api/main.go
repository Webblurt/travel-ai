package main

import (
	"fmt"
	"net/http"
	handlers "travel-ai/internal/http"
)

func main() {
	mux := http.NewServeMux()
	handlers.SetupRoutes(mux)

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	fmt.Println("Backend running on :8080")
	http.ListenAndServe(":8080", mux)
}
