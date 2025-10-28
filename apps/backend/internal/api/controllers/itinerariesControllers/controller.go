package itinerariescontrollers

import services "travel-ai/internal/services"

type ItinerariesController struct {
	Service services.ServiceInterface
}

func NewItinerariesController(service services.ServiceInterface) *ItinerariesController {
	return &ItinerariesController{
		Service: service,
	}
}
