package models

type GetItinerariesReq struct {
	PerPage  string `json:"per_page"`
	Page     string `json:"page"`
	City     string `json:"city"`
	Days     string `json:"days"`
	Budget   string `json:"budget"`
	Currency string `json:"currency"`
	UserID   string
}

type GetItineraryReq struct {
	City     string `json:"city"`
	Days     string `json:"days"`
	Budget   string `json:"budget"`
	Currency string `json:"currency"`
	UserID   string
}
