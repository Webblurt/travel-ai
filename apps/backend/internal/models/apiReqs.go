package models

type GetItineraryReq struct {
	City     string `json:"city"`
	Days     string `json:"days"`
	Budget   string `json:"budget"`
	Currency string `json:"currency"`
}
