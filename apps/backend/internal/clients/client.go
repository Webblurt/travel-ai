package clients

import (
	"errors"
	utils "travel-ai/internal/utils"

	"github.com/sashabaranov/go-openai"
)

type ClientsInterface interface {
	AskOpenAI(promptType string, req interface{}) (interface{}, error)
}

type Client struct {
	Name   string
	URL    string
	Key    string
	log    *utils.Logger
	cfg    *utils.Config
	openai *openai.Client
}

func NewClient(name, url, key string, log *utils.Logger, cfg *utils.Config) *Client {
	return &Client{
		Name:   name,
		URL:    url,
		Key:    key,
		log:    log,
		cfg:    cfg,
		openai: openai.NewClient(key),
	}
}

func CreateClients(cfg *utils.Config, log *utils.Logger) ([]*Client, error) {
	var clients []*Client

	for _, api := range cfg.APIS {
		if api.Name == "" || api.URL == "" {
			log.Warn("api configuration missing something")
			continue
		}
		client := NewClient(api.Name, api.URL, api.Key, log, cfg)
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return nil, errors.New("0 clients created")
	}

	return clients, nil
}
