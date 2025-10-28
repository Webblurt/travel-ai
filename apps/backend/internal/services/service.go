package services

import (
	"errors"
	clients "travel-ai/internal/clients"
	models "travel-ai/internal/models"
	utils "travel-ai/internal/utils"
)

type ServiceInterface interface {
	Validate(token string) (string, error)
	GetItinerary(filters models.GetItineraryReq) (models.GetItineraryResp, error)
}

type Service struct {
	clients    []*clients.Client
	components *utils.Components
	log        *utils.Logger
	cfg        *utils.Config
}

func NewService(cfg *utils.Config, clientsList []*clients.Client, log *utils.Logger) (*Service, error) {
	if len(clientsList) == 0 {
		return nil, errors.New("no clients provided")
	}

	return &Service{
		clients:    clientsList,
		components: &cfg.Components,
		log:        log,
		cfg:        cfg,
	}, nil
}

func (s *Service) SelectClientByName(name string) (*clients.Client, error) {
	for _, client := range s.clients {
		if client.Name == name {
			return client, nil
		}
	}
	return nil, errors.New("client not found")
}
