package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found, using system environment variables", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is missing in .env")
	}

	client := openai.NewClient(apiKey)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/itinerary", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		type Request struct {
			City string `json:"city"`
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
			return
		}

		if req.City == "" {
			http.Error(w, `{"error":"city is required"}`, http.StatusBadRequest)
			return
		}

		prompt := fmt.Sprintf(`Составь подробный 3-дневный туристический маршрут по городу %s. 
Предложи достопримечательности, кафе и местные активности. Пиши кратко и дружелюбно.`, req.City)

		resp, err := client.CreateChatCompletion(r.Context(), openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "Ты — AI-гид по путешествиям."},
				{Role: "user", Content: prompt},
			},
		})
		if err != nil {
			log.Println("OpenAI error:", err)
			http.Error(w, `{"error":"AI request failed"}`, http.StatusInternalServerError)
			return
		}

		answer := resp.Choices[0].Message.Content

		json.NewEncoder(w).Encode(map[string]string{
			"itinerary": answer,
		})
	})

	// static frontend files
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)

	fmt.Println("Backend running on :8080")
	http.ListenAndServe(":8080", mux)
}
