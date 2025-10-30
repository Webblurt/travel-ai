package services

import (
	"context"
	"errors"
	clients "travel-ai/internal/clients"
	models "travel-ai/internal/models"
	repositories "travel-ai/internal/repositories"
	utils "travel-ai/internal/utils"
)

type ServiceInterface interface {
	Validate(token string) (string, error)
	GetItinerary(filters models.GetItineraryReq) (models.GetItineraryResp, error)
	GetExactItinerary(id string) (models.ExactIniterary, error)
	GetExactDayItinerary(id, dayNum string) (models.ExactDayPlan, error)
	GetTimeOfDayItinerary(id, dayNum, timeOfDay, userID string) (models.ExactItemPlan, error)
}

type Service struct {
	clients    []*clients.Client
	repository *repositories.Repository
	components *utils.Components
	log        *utils.Logger
	cfg        *utils.Config
	ctx        context.Context
}

func NewService(cfg *utils.Config, clientsList []*clients.Client, repo *repositories.Repository, log *utils.Logger) (*Service, error) {
	if len(clientsList) == 0 {
		return nil, errors.New("no clients provided")
	}

	ctx := context.Background()

	return &Service{
		clients:    clientsList,
		repository: repo,
		components: &cfg.Components,
		log:        log,
		cfg:        cfg,
		ctx:        ctx,
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
