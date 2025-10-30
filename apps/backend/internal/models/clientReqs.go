package models

type OpenAIGenerateItineraryReq struct {
	City     string
	Days     string
	Budget   string
	Currency string
}

type OpenAIGenerateItineraryDetailsReq struct {
	TimeOfDay    string
	Place        string
	Description  string
	ActivityType string
	Cost         string
	Duration     string
	Address      string
	City         string
}
