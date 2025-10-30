package itinerariescontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	services "travel-ai/internal/services"
)

type ItinerariesController struct {
	Service services.ServiceInterface
}

func NewItinerariesController(service services.ServiceInterface) *ItinerariesController {
	return &ItinerariesController{
		Service: service,
	}
}

func fetchAuthorizationHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", errors.New("invalid authorization header format")
	}
	token := authHeader[len(bearerPrefix):]

	return token, nil
}

func (ic *ItinerariesController) withAuth(handler func(w http.ResponseWriter, r *http.Request, token string, id string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := fetchAuthorizationHeader(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		id, err := ic.Service.Validate(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		handler(w, r, token, id)
	}
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
